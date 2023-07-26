package main

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/db"
	migrations "KirsanovStavkaTV/migrations"
	"KirsanovStavkaTV/server"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	fmt.Println(os.Getenv("DB_TYPE"))
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
	migrations.Migrate()
	server.NewServer(service)
}
