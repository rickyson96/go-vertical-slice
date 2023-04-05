package server_test

import (
	"context"

	"github.com/jackc/pgx/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rickyson96/go-vertical-slice/internal/domain/db"
	"github.com/rickyson96/go-vertical-slice/internal/server"
)

var _ = Describe("Database", func() {
	Context("Connecting to database", func() {
		It("should be able to connect to database", func() {
			connPool, err := server.Conn(nil)
			Expect(err).NotTo(HaveOccurred())

			var result int
			connPool.QueryRow(context.Background(), "SELECT 1").Scan(&result)

			Expect(result).To(Equal(1))
		})

		It("should be able to use existing connection", func() {
			connPool, err := server.Conn(nil)
			Expect(err).NotTo(HaveOccurred())

			newConn, err := server.Conn(nil)
			Expect(err).NotTo(HaveOccurred())

			Expect(newConn).To(Equal(connPool))
		})
	})

	It("migrating to newest schema", func() {

		ctx := context.Background()

		conn, err := server.Conn(nil)
		Expect(err).NotTo(HaveOccurred())

		server.MigrateSchema(ctx, conn)

		_, err = db.New(conn).GetConfig(ctx, "test")
		Expect(err).To(MatchError(pgx.ErrNoRows))

	})

})
