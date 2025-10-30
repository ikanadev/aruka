package utils

import "os"

type Config struct {
	GrpcPort         string
	DBConn           string
	MigrationsSource string
	AnthropicAPIKey  string
}

var _config = Config{
	GrpcPort:         os.Getenv("GRPC_PORT"),
	DBConn:           os.Getenv("DATABASE"),
	MigrationsSource: os.Getenv("MIGRATIONS_SOURCE"),
	AnthropicAPIKey:  os.Getenv("ANTHROPIC_API_KEY"),
}

func GetConfig() Config {
	return _config
}
