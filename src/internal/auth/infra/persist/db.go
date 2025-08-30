package persist

import (
	"context"
	"log"
	"time"

	"github.com/college-redberry/open-adoption/internal/auth/infra/constants"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, constants.DB_URL)
	if err != nil {
		panic("error while creating db pool: " + err.Error())
	}

	// Check connectivity
	if err := dbpool.Ping(ctx); err != nil {
		panic("error pinging db: " + err.Error())
	}

	log.Println("Database connection established with pgxpool.")
	return dbpool
}
