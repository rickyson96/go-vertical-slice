package server

import (
	"context"
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/tern/migrate"
	"github.com/rickyson96/go-vertical-slice/internal/server/config"
	"github.com/spf13/viper"
)

var pool *pgxpool.Pool

func Conn(ctx context.Context) (*pgxpool.Pool, error) {
	var err error

	if pool == nil {
		dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			viper.GetString(config.DatabaseUser),
			viper.GetString(config.DatabasePassword),
			viper.GetString(config.DatabaseHost),
			viper.GetString(config.DatabasePort),
			viper.GetString(config.DatabaseName),
		)

		if ctx == nil {
			ctx = context.Background()
		}
		pool, err = pgxpool.Connect(ctx, dbURL)
	}

	return pool, err
}

func MigrateSchema(ctx context.Context, pool *pgxpool.Pool) error {
	return pool.AcquireFunc(ctx, func(c *pgxpool.Conn) error {
		m, err := migrate.NewMigrator(ctx, c.Conn(), "public.schema_version")
		if err != nil {
			return err
		}

		_, thisFile, _, _ := runtime.Caller(0)
		mpath := filepath.Join(thisFile, "../../../sql/schema")

		err = m.LoadMigrations(mpath)
		if err != nil {
			return err
		}

		err = m.Migrate(ctx)
		return err
	})
}
