package main

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/db"
	"KirsanovStavkaTV/server"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

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
}
