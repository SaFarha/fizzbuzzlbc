package configuration

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"strings"
)

type Configuration struct {
	Env string

	Port string

	PostgresqlHost     string
	PostgresqlPort     int
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbName   string
	PostgresqlSsl      string
}

func LoadConfig() Configuration {
	_ = godotenv.Load(".env")
	conf := Configuration{}

	viper.AutomaticEnv()

	conf.Env = viper.GetString("env")

	// overwride config with env var API_<MY>_<VAR>
	viper.SetEnvPrefix("api")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// define default values
	viper.SetDefault("port", 8080)

	viper.SetDefault("postgresql.host", "localhost")
	viper.SetDefault("postgresql.port", 5433)
	viper.SetDefault("postgresql.user", "user")
	viper.SetDefault("postgresql.password", "password")
	viper.SetDefault("postgresql.database", "fizzbuzzlbc-api-test")
	viper.SetDefault("postgresql.ssl", "disable")

	// load configuration
	conf.Port = viper.GetString("port")
	conf.PostgresqlHost = viper.GetString("postgresql.host")
	conf.PostgresqlPort = viper.GetInt("postgresql.port")
	conf.PostgresqlUser = viper.GetString("postgresql.user")
	conf.PostgresqlPassword = viper.GetString("postgresql.password")
	conf.PostgresqlDbName = viper.GetString("postgresql.database")
	conf.PostgresqlSsl = viper.GetString("postgresql.ssl")

	return conf
}
