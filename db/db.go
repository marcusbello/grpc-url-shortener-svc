package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teris-io/shortid"
	"grpc-url-shortener-svc/model"
	"log"
	"time"
)

var ErrNotFound = fmt.Errorf("missing record")

type PGData struct {
	client  *pgxpool.Pool
	timeout time.Duration
}

func newPGDB(dsn string) (*pgxpool.Pool, error) {
	ctx := context.Background()
	dbpool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	err = dbpool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping DB connection: %v\n", err)
		return nil, err
	}

	return dbpool, nil

}

func NewPGRepo(dsn string) (*PGData, error) {
	repo := &PGData{}
	client, err := newPGDB(dsn)
	if err != nil {
		log.Printf("Unable to create connection database: %v\n", err)
		return nil, err
	}

	repo.client = client
	return repo, nil
}

func (p *PGData) Add(ctx context.Context, r *model.AddRequest) error {

	stmt := `INSERT INTO shorturl (url, short_code, created_by, created_at) VALUES($1, $2, $3, $4) RETURNING *;`

	_, err := p.client.Exec(ctx, stmt, r.Url, r.ShortCode, r.CreatedBy, r.CreatedAt)
	if err != nil {
		log.Printf("Failed to insert data into rows: %s", err)
		return err
	}

	return nil
}

func (p *PGData) Show(ctx context.Context, id string, data interface{}) error {
	stmt := `select id, email, password from users where id=$1`
	rowScan, err := p.client.Query(ctx, stmt, id)
	if err != nil {
		return ErrNotFound
	}

	return rowScan.Scan(data)

}

func (p *PGData) GenShortCode() (string, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2341)
	if err != nil {
		return "", err
	}
	return sid.MustGenerate(), nil
}
