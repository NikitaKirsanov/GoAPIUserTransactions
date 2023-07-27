package db

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type RedisProvider struct {
	DB *redis.Client
}

func (r RedisProvider) Provide() contracts.DatabaseProvider {
	fmt.Println(os.Getenv("REDIS-DB"))
	redisAddr := os.Getenv("REDIS-ADDR")
	redisPassword := os.Getenv("REDIS-PASSWORD")
	redisDB, err := strconv.Atoi(os.Getenv("REDIS-DB"))
	if err != nil {
		panic("could't convert reddis DB to type int")
	}
	r.DB = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	return r
}

func (r RedisProvider) FindUser(id int) models.User {
	idStr := constants.RedisUserPrefix + strconv.Itoa(id)
	user := models.User{}
	userRecord, err := r.DB.Get(context.Background(), idStr).Result()
	if err != nil {
		return user
	}

	json.Unmarshal([]byte(userRecord), &user)
	return user
}

func (r RedisProvider) GetUsers() []models.User {
	users := []models.User{}
	redisResult, err := r.DB.Keys(context.Background(), constants.RedisUserPrefix+"*").Result()
	if err != nil {
		panic(err)
	}
	for _, key := range redisResult {
		user := models.User{}
		json.Unmarshal([]byte(key), &user)
		users = append(users, user)
	}

	return users
}

func (r RedisProvider) MakeTransfer(t *models.Transaction) error {
	ctx := context.Background()
	userFromString, err := r.DB.Get(ctx, constants.RedisUserPrefix+strconv.Itoa(t.UserFrom)).Result()
	if err != nil {
		return err
	}
	userFrom := models.User{}
	err = json.Unmarshal([]byte(userFromString), &userFrom)
	if err != nil {
		return err
	}

	userToString, err := r.DB.Get(ctx, constants.RedisUserPrefix+strconv.Itoa(t.UserFrom)).Result()
	if err != nil {
		return err
	}
	userTo := models.User{}
	err = json.Unmarshal([]byte(userToString), &userTo)
	if err != nil {
		return err
	}

	userFrom.Balance = userFrom.GetBalance() - t.Amount
	userTo.Balance = userTo.GetBalance() + t.Amount

	transactionRecords, err := r.DB.Keys(ctx, constants.RedisTransactionPrifix+"*").Result()
	if err != nil {
		return err
	}

	transactionId := 1
	len := len(transactionRecords)
	if len > 0 {
		lastTransaction := models.Transaction{}
		lastTransactionRec := transactionRecords[len-1]
		err = json.Unmarshal([]byte(lastTransactionRec), &lastTransaction)
		if err != nil {
			return err
		}
		transactionId = lastTransaction.Id + 1
	}

	t.Id = transactionId

	userFromMarshalled, err := json.Marshal(userFrom)
	if err != nil {
		return err
	}
	userToMarshalled, err := json.Marshal(userTo)
	if err != nil {
		return err
	}
	transactionMarshalled, err := json.Marshal(t)
	if err != nil {
		return err
	}
	err = r.DB.Watch(ctx, func(tx *redis.Tx) error {
		transactionExists, err := tx.Exists(ctx, constants.RedisTransactionPrifix+fmt.Sprint(t.Id)).Result()
		if err != nil {
			return err
		}
		if transactionExists != 0 {
			return errors.New("Transaction already exists, contact N.Kirsanov")
		}
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			if _, err = pipe.Del(
				ctx,
				constants.RedisUserPrefix+fmt.Sprint(userFrom.Id),
				constants.RedisUserPrefix+fmt.Sprint(userTo.Id),
			).Result(); err != nil {
				return err
			}

			if _, err = pipe.Set(
				ctx,
				constants.RedisUserPrefix+fmt.Sprint(userFrom.Id),
				string(userFromMarshalled),
				time.Duration(time.Duration.Hours(24)),
			).Result(); err != nil {
				return err
			}

			if _, err = pipe.Set(
				ctx,
				constants.RedisUserPrefix+fmt.Sprint(userTo.Id),
				string(userToMarshalled),
				time.Duration(time.Duration.Hours(24)),
			).Result(); err != nil {
				return err
			}

			if _, err = pipe.Set(
				ctx,
				constants.RedisTransactionPrifix+fmt.Sprint(t.Id),
				string(transactionMarshalled),
				time.Duration(time.Duration.Hours(24)),
			).Result(); err != nil {
				return err
			}

			return nil
		})

		return err
	})
	if err != nil {
		return errors.New("Saving transaction failed, try again later")
	}

	return nil
}
