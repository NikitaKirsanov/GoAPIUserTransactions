package constants

const (
	AppName                = "stavkatv"
	SchemaName             = "stavkatv"
	MigrationsTableName    = SchemaName + ".migrations"
	UsersTableName         = SchemaName + ".users"
	TransactionsTableName  = SchemaName + ".transactions"
	DBTypePostgres         = "POSTGRES"
	DBTypeRedis            = "REDIS"
	RedisUserPrefix        = "user_"
	RedisTransactionPrifix = "transaction_"
)
