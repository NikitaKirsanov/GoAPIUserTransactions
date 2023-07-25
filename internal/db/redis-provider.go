package db

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/models"
	"context"
	"encoding/json"
	"os"
	"strconv"

	redis "github.com/redis/go-redis/v9"
)

type RedisProvider struct {
	DB *redis.Client
}

func (r RedisProvider) Provide() {
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

	userFrom.SetBalance(userFrom.GetBalance() - t.Amount)
	userTo.SetBalance(userTo.GetBalance() + t.Amount)

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

	/**
	*
	*
	* Продолжить отсюда
	*
	*
	err := a.rds.Watch(ctx, func(tx *redis.Tx) error {
		// You can run more commands here. You will use `tx` though.
		// e.g. tx.HGET()
		// Note: `tx` is not part of transactional `pipe` below so any "SET" operation will be independent.

		_, err := tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			// 1- Set holder
			if _, err := pipe.HSetNX(
				ctx,
				account.CacheHashRootKey(holder.ID),
				account.CacheHashHolderField(),
				&holder,
			).Result(); err != nil {
				return fmt.Errorf("create: holder: %w", err)
			}
			//tx.Expire()
			//

			// 2- Set accounts within holder
			for _, acc := range accounts {
				if _, err := pipe.HSetNX(
					ctx,
					account.CacheHashRootKey(holder.ID),
					account.CacheHashAccountField(acc.Type),
					acc,
				).Result(); err != nil {
					return fmt.Errorf("create: account: %w", err)
				}
				//pipe.Expire()
			}
			//

			return nil
		})

		return err
	})
	if err != nil {
		return fmt.Errorf("create: transaction: %w", err)
	}
	*/

	return nil
}
