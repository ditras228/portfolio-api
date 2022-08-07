package postgres

import (
	"context"
	"fmt"
	"log"
	config "portfolio-api/internal"
	"portfolio-api/pkg/utils"
	"time"
)
import "github.com/jackc/pgconn"
import "github.com/jackc/pgx/v4"
import "github.com/jackc/pgx/v4/pgxpool"

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, sc config.StorageConfig) (pool *pgxpool.Pool, err error) {
	dsc := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsc)
		if err != nil {
			return err
		}

		return nil

	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error with tries postgres")
	}
	return pool, nil
}
