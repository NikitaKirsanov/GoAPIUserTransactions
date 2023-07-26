package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName)

	if _, err := sql.Open("postgres", url); err != nil {
		panic(fmt.Errorf("Postgres connect error : (%v)", err))
	} else {
		panic("THIS SHIT WORKS")
	}

	/*
			var dbProvider contracts.DatabaseProvider
			switch conn := os.Getenv("DB_TYPE"); conn {
			case constants.DBTypePostgres:
				dbProvider = &db.PostgresProvider{}
			case constants.DBTypeRedis:
				dbProvider = &db.RedisProvider{}
			default:
				panic("unknown db type")
			}
			dbProvider.Provide()


		service := server.NewService(dbProvider)
		//migrations.Migrate()
		server.NewServer(service)
	*/
}
