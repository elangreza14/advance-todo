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
	// DbSQLOption is option to use in app
	DbSQLOption struct {
		MaxLifeTime  time.Duration
		MaxIdleConns int
		MaxOpenConns int
	}

	// migrationLogger is for logger go-migrate
	migrationLogger struct{}
)

func (c DbSQLOption) apply(config *Configuration) error {
	connString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		config.Env.PostgresUser,
		config.Env.PostgresPassword,
		config.Env.PostgresHostname,
		config.Env.PostgresPort,
		config.Env.PostgresDB,
		config.Env.PostgresSsl)

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		return err
	}

	// migrate db from migration
	migration, err := migrate.New(fmt.Sprintf("file://%v", config.Env.PostgresMigrationFolder), connString)
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
	config.DbSQL = pool
	return nil
}

func (m migrationLogger) Printf(format string, v ...interface{}) {
	formatter := fmt.Sprintf("migration | %v", format)
	log.Printf(formatter, v...)
}

func (m migrationLogger) Verbose() bool {
	return true
}

// WithDBSql is option interface to use sql database in app
func WithDBSql(c DbSQLOption) Option {
	return c
}
