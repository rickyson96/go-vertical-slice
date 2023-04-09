package server

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/rickyson96/go-vertical-slice/internal/server/config"
	"github.com/spf13/viper"
)

func CleanUpDatabase(ctx context.Context, conn *pgxpool.Pool) error {
	_, err := conn.Exec(ctx, `DROP SCHEMA public CASCADE; CREATE SCHEMA public;`)
	if err != nil {
		return err
	}

	return MigrateSchema(ctx, conn)
}

func SetupSuiteTest() {
	var container *gnomock.Container
	ctx := context.Background()

	BeforeSuite(func() {
		var err error

		dbName := "testdb"
		dbUser := "testuser"
		dbPassword := "testpassword"

		p := postgres.Preset(postgres.WithDatabase(dbName), postgres.WithUser(dbUser, dbPassword))
		container, err = gnomock.Start(p)
		Expect(err).NotTo(HaveOccurred())

		viper.Set(config.DatabaseUser, dbUser)
		viper.Set(config.DatabasePassword, dbPassword)
		viper.Set(config.DatabaseHost, container.Host)
		viper.Set(config.DatabasePort, container.DefaultPort())
		viper.Set(config.DatabaseName, dbName)

		conn, err := Conn(ctx)
		Expect(err).NotTo(HaveOccurred())

		err = MigrateSchema(ctx, conn)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterSuite(func() {
		gnomock.Stop(container)
	})
}

func BeforeSuite(func func()) {
	panic("unimplemented")
}
