package config

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type (
	DbSqlOption struct {
		MaxLifeTime  time.Duration
		MaxIdleConns int
		MaxOpenConns int
	}

	// this is for logger go-migrate
	migrationLogger struct{}
)

func (c DbSqlOption) apply(config *Configuration) error {
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		config.Env.POSTGRES_USER,
		config.Env.POSTGRES_PASSWORD,
		config.Env.POSTGRES_HOSTNAME,
		config.Env.POSTGRES_PORT,
		config.Env.POSTGRES_DB,
		config.Env.POSTGRES_SSL)

	pool, err := sql.Open("postgres", connString)
	if err != nil {
		return err
	}

	// set limiter configuration
	pool.SetConnMaxLifetime(c.MaxLifeTime)
	pool.SetMaxIdleConns(c.MaxIdleConns)
	pool.SetMaxOpenConns(c.MaxOpenConns)

	// if ping exceeding 5 second
	// will return
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		return err
	}

	// migrate db from migration
	migration, err := migrate.New(fmt.Sprintf("file://%v", config.Env.POSTGRES_MIGRATION_FOLDER), connString)
	if err != nil {
		return err
	}

	migration.Log = new(migrationLogger)

	// apply migration to DB
	if err := migration.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}

	// assign pool into postgres so app can use it
	config.DbSql = pool
	return nil
}

func (m migrationLogger) Printf(format string, v ...interface{}) {
	formatter := fmt.Sprintf("migration | %v", format)
	log.Printf(formatter, v...)
}

func (m migrationLogger) Verbose() bool {
	return true
}

// later we can also adding mysql etc
func WithPostgres(c DbSqlOption) Option {
	return DbSqlOption(c)
}
