package postgres

import (
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type Storage struct {
	db *pgdb.DB
}

func NewPostgres(db *pgdb.DB) *Storage {
	return &Storage{db: db}
}

func (s *Storage) SaveURL(urlToSave string, alias string) error {

	query := fmt.Sprintf("INSERT INTO link(url, alias) VALUES($1, $2) RETURNING id")

	if err := s.db.ExecRaw(query, urlToSave, alias); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetURL(alias string) ([]string, error) {

	query := fmt.Sprintf("SELECT url FROM link WHERE alias = $1")

	var recievedURL []string

	if err := s.db.SelectRaw(&recievedURL, query, alias); err != nil {
		return nil, err
	}

	return recievedURL, nil
}
