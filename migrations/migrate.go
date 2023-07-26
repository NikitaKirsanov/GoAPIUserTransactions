package migrations

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/golang-migrate/migrate"
	_ "github.com/lib/pq"
	redis "github.com/redis/go-redis/v9"
)

func Migrate() {
	switch conn := os.Getenv("DB_TYPE"); conn {
	case constants.DBTypePostgres:
		dbHost := os.Getenv("POSTGRES_HOST")
		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbPort := os.Getenv("POSTGRES_PORT")
		dbName := os.Getenv("POSTGRES_DB")
		url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			dbUser,
			dbPassword,
			dbHost,
			dbPort,
			dbName)
		m, err := migrate.New(
			"file:///migrations",
			url)
		if err != nil {
			panic("Couldn't migrate users")
		}
		m.Steps(2)

	case constants.DBTypeRedis:
		randBalanceOne := uint(rand.Uint64())
		randBalanceTwo := uint(rand.Uint64())
		createdAt := time.Now()
		userOne := models.User{
			Id:        1,
			Balance:   randBalanceOne,
			CreatedAt: &createdAt,
		}
		userTwo := models.User{
			Id:        2,
			Balance:   randBalanceTwo,
			CreatedAt: &createdAt,
		}
		redisAddr := os.Getenv("REDIS-ADDR")
		redisPassword := os.Getenv("REDIS-PASSWORD")
		redisDB, err := strconv.Atoi(os.Getenv("REDIS-DB"))
		if err != nil {
			panic("could't convert reddis DB to type int")
		}
		DB := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDB,
		})
		userOneJson, err := json.Marshal(userOne)
		if err != nil {
			panic("couldn't migrate users")
		}
		resOne, err := DB.Set(context.Background(), constants.RedisUserPrefix+fmt.Sprint(userOne.Id), string(userOneJson), time.Duration(time.Duration.Hours(24))).Result()
		if err != nil {
			panic("couldn't migrate users")
		}
		//Здесь должен быть какой-нибудь логгер, но пока просто пишем в stdOut
		fmt.Println(fmt.Sprintf("Added 1 user during migration user:%s", resOne))

		userTwoJson, err := json.Marshal(userTwo)
		if err != nil {
			panic("couldn't migrate users")
		}
		resTwo, err := DB.Set(context.Background(), constants.RedisUserPrefix+fmt.Sprint(userTwo.Id), string(userTwoJson), time.Duration(time.Duration.Hours(24))).Result()
		if err != nil {
			panic("couldn't migrate users")
		}
		fmt.Println(fmt.Sprintf("Added 1 user during migration user:%s", resTwo))
	default:
		panic("unknown db type")
	}
}
