package createuser

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/julienschmidt/httprouter"
	"github.com/rickyson96/go-vertical-slice/internal/domain/db"
)

type handler struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) httprouter.Handle {
	return handler{
		pool: pool,
	}.Handle
}

type Request struct {
	Name string `json:"name"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request Request
	ctx := r.Context()

	tx, _ := h.pool.Begin(ctx)
	defer tx.Rollback(ctx)

	json.NewDecoder(r.Body).Decode(&request)
	defer tx.Rollback(ctx)

	txQueries := db.New(h.pool).WithTx(tx)
	err := txQueries.CreateUser(ctx, request.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error"))
		return
	}
	tx.Commit(ctx)

	w.WriteHeader(http.StatusCreated)
}
