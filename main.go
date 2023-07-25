package main

import (
	"KirsanovStavkaTV/internal/constants"
	"KirsanovStavkaTV/internal/contracts"
	"KirsanovStavkaTV/internal/db"
	"os"
)

func main() {
	var dbProvider contracts.DatabaseProvider
	switch conn := os.Getenv("DB-TYPE"); conn {
	case constants.DBTypePostgres:
		dbProvider = &db.PostgresProvider{}
	case constants.DBTypeRedis:
		dbProvider = &db.RedisProvider{}
	default:
		panic("uncnown db type")
	}
	dbProvider.Provide()

}
