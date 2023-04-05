package server_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/postgres"
	"github.com/spf13/viper"
	"github.com/rickyson96/go-vertical-slice/internal/server/config"
)

func TestServer(t *testing.T) {
	var container *gnomock.Container

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
	})

	AfterSuite(func() {
		gnomock.Stop(container)
	})

	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}
