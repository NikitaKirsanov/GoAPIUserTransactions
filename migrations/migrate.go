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

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	redis "github.com/redis/go-redis/v9"
)

func Migrate() {
	switch conn := os.Getenv("DB_TYPE"); conn {
	case constants.DBTypePostgres:
		break
	case constants.DBTypeRedis:
		randBalanceOne := uint(rand.Uint64())
		randBalanceTwo := uint(rand.Uint64())
		userOne := models.User{
			Id:      1,
			Balance: randBalanceOne,
		}
		userTwo := models.User{
			Id:      2,
			Balance: randBalanceTwo,
		}
		redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err != nil {
			panic(fmt.Sprintf("could't convert redis DB to type int %s, err: %s", os.Getenv("REDIS_DB"), err))
		}
		DB := redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
			DB:       redisDB,
		})
		userOneJson, err := json.Marshal(userOne)
		if err != nil {
			panic("couldn't migrate users")
		}
		resOne, err := DB.Set(context.Background(), constants.RedisUserPrefix+fmt.Sprint(userOne.Id), string(userOneJson), time.Duration(time.Duration.Hours(24))).Result()
		if err != nil {
			panic(fmt.Sprintf("couldn't migrate users err:%s", err))
		}
		fmt.Println(fmt.Sprintf("userOne:%s", string(userOneJson)))
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
		fmt.Println(fmt.Sprintf("userOne:%s", string(userTwoJson)))
		fmt.Println(fmt.Sprintf("Added 1 user during migration user:%s", resTwo))
	default:
		panic("unknown db type")
	}
}
