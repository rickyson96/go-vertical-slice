package server_test

import (
	"context"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rickyson96/go-vertical-slice/internal/domain/db"
	"github.com/rickyson96/go-vertical-slice/internal/server"
)

var _ = Describe("Dbtesthelper", func() {
	ctx := context.Background()
	It("should be able to clean up database", func() {
		conn, err := server.Conn(ctx)
		Expect(err).NotTo(HaveOccurred())

		Expect(server.MigrateSchema(ctx, conn)).To(Succeed())

		_, err = conn.Exec(ctx, `INSERT INTO configs VALUES ('test','"test"');`)
		Expect(err).NotTo(HaveOccurred())

		res, err := db.New(conn).GetConfig(ctx, "test")
		Expect(err).NotTo(HaveOccurred())
		value := &pgtype.JSONB{}
		value.Set(`"test"`)
		Expect(res).To(Equal(db.Config{
			Key:   "test",
			Value: *value,
		}))

		server.CleanUpDatabase(ctx, conn)

		_, err = db.New(conn).GetConfig(ctx, "test")
		Expect(err).To(MatchError(pgx.ErrNoRows))
	})
})
