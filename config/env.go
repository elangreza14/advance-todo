package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type (
	// Env is list all of env
	Env struct {
		PostgresHostname        string `mapstructure:"POSTGRES_HOSTNAME"`
		PostgresSsl             string `mapstructure:"POSTGRES_SSL"`
		PostgresUser            string `mapstructure:"POSTGRES_USER"`
		PostgresPassword        string `mapstructure:"POSTGRES_PASSWORD"`
		PostgresDB              string `mapstructure:"POSTGRES_DB"`
		PostgresPort            int32  `mapstructure:"POSTGRES_PORT"`
		PostgresMigrationFolder string `mapstructure:"POSTGRES_MIGRATION_FOLDER"`
		RedisHostName           string `mapstructure:"REDIS_HOSTNAME"`
		RedisPass               string `mapstructure:"REDIS_PASS"`
		RedisPort               int32  `mapstructure:"REDIS_PORT"`
		RedisDB                 int    `mapstructure:"REDIS_DB"`
		RedisReplicationMode    string `mapstructure:"REDIS_REPLICATION_MODE"`
		TokenKey                string `mapstructure:"TOKEN_KEY"`
		XApiKey                 string `mapstructure:"X_API_KEY"`
	}
)

// NewEnv is constructor for getting env cross platform
// can be use with direct env or env file
func NewEnv() (*Env, error) {
	env := &Env{}
	envBase := os.Getenv("MODE")

	switch {
	case envBase != "":
		viper.AddConfigPath(".")
		viper.SetConfigName(envBase)
		viper.SetConfigType("env")
		viper.AutomaticEnv()

		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}

		err = viper.Unmarshal(&env)
		if err != nil {
			return nil, err
		}

		return env, nil
	default:
		// postgres
		env.PostgresHostname = env.getString("POSTGRES_HOSTNAME")
		env.PostgresSsl = env.getString("POSTGRES_SSL")
		env.PostgresUser = env.getString("POSTGRES_USER")
		env.PostgresPassword = env.getString("POSTGRES_PASSWORD")
		env.PostgresDB = env.getString("POSTGRES_DB")
		pgPort, err := env.getInt32("POSTGRES_PORT")
		if err != nil {
			return nil, err
		}
		env.PostgresPort = *pgPort
		env.PostgresMigrationFolder = env.getString("POSTGRES_MIGRATION_FOLDER")

		// redis
		env.RedisHostName = env.getString("REDIS_HOSTNAME")
		env.RedisPass = env.getString("REDIS_PASS")
		rdPort, err := env.getInt32("REDIS_PORT")
		if err != nil {
			return nil, err
		}
		env.RedisPort = *rdPort
		rdDB, err := env.getInt("REDIS_DB")
		if err != nil {
			return nil, err
		}
		env.RedisDB = *rdDB
		env.RedisReplicationMode = env.getString("REDIS_REPLICATION_MODE")

		// jwt
		env.TokenKey = env.getString("TOKEN_KEY")

		// x-api-key
		env.XApiKey = env.getString("x_API_KEY")

		return env, nil
	}
}

func (e *Env) getString(envName string) string {
	if res, ok := os.LookupEnv(envName); ok {
		return res
	}
	return ""
}

func (e *Env) getInt32(envName string) (*int32, error) {
	if res, ok := os.LookupEnv(envName); ok {
		resInt, err := strconv.Atoi(res)
		if err != nil {
			return nil, err
		}

		resInt32 := int32(resInt)
		return &resInt32, nil
	}

	return nil, errors.New("error parsing to int 32 data")
}

func (e *Env) getInt(envName string) (*int, error) {
	if res, ok := os.LookupEnv(envName); ok {
		resInt, err := strconv.Atoi(res)
		if err != nil {
			return nil, err
		}

		return &resInt, nil
	}

	return nil, errors.New("error parsing to int data")
}
