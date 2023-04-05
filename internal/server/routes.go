package server

import (
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rickyson96/go-vertical-slice/internal/features/healthcheck"
)

func Routes(pool *pgxpool.Pool) http.Handler {
	router := httprouter.New()

	router.GET("/health", healthcheck.New())

	return router
}
