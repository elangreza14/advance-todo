package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type (
	Env struct {
		POSTGRES_HOSTNAME         string `mapstructure:"POSTGRES_HOSTNAME"`
		POSTGRES_SSL              string `mapstructure:"POSTGRES_SSL"`
		POSTGRES_USER             string `mapstructure:"POSTGRES_USER"`
		POSTGRES_PASSWORD         string `mapstructure:"POSTGRES_PASSWORD"`
		POSTGRES_DB               string `mapstructure:"POSTGRES_DB"`
		POSTGRES_PORT             int32  `mapstructure:"POSTGRES_PORT"`
		POSTGRES_MIGRATION_FOLDER string `mapstructure:"POSTGRES_MIGRATION_FOLDER"`
		REDIS_HOSTNAME            string `mapstructure:"REDIS_HOSTNAME"`
		REDIS_PASS                string `mapstructure:"REDIS_PASS"`
		REDIS_PORT                int32  `mapstructure:"REDIS_PORT"`
		REDIS_DB                  int    `mapstructure:"REDIS_DB"`
		REDIS_REPLICATION_MODE    string `mapstructure:"REDIS_REPLICATION_MODE"`
		TOKEN_KEY                 string `mapstructure:"TOKEN_KEY"`
	}
)

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
		env.POSTGRES_HOSTNAME = env.getString("POSTGRES_HOSTNAME")
		env.POSTGRES_SSL = env.getString("POSTGRES_SSL")
		env.POSTGRES_USER = env.getString("POSTGRES_USER")
		env.POSTGRES_PASSWORD = env.getString("POSTGRES_PASSWORD")
		env.POSTGRES_DB = env.getString("POSTGRES_DB")
		pgPort, err := env.getInt32("POSTGRES_PORT")
		if err != nil {
			return nil, err
		}
		env.POSTGRES_PORT = *pgPort
		env.POSTGRES_MIGRATION_FOLDER = env.getString("POSTGRES_MIGRATION_FOLDER")

		// redis
		env.REDIS_HOSTNAME = env.getString("REDIS_HOSTNAME")
		env.REDIS_PASS = env.getString("REDIS_PASS")
		rdPort, err := env.getInt32("REDIS_PORT")
		if err != nil {
			return nil, err
		}
		env.REDIS_PORT = *rdPort
		rdDB, err := env.getInt("REDIS_DB")
		if err != nil {
			return nil, err
		}
		env.REDIS_DB = *rdDB
		env.REDIS_REPLICATION_MODE = env.getString("REDIS_REPLICATION_MODE")

		// jwt
		env.TOKEN_KEY = env.getString("TOKEN_KEY")

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
