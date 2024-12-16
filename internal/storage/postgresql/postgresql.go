package postgresql

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(host, port, user, password, dbname, sslmode string) (*Storage, error) {
	const op = "storage.sqlite.New"
	connString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Stop() error {
	return s.db.Close()
}

func (s *Storage) CreateTransaction(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.sqlite.CreateTransaction"

	stmt, err := s.db.Prepare("INSERT INTO transactions(email, pass_hash) VALUES($1, $2) RETURNING id")
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.ExecContext(ctx, email, passHash)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
