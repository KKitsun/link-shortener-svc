package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/KKitsun/link-shortener-svc/internal/storage"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const linkTableName = "link"

func newLinkQ(db *pgdb.DB) storage.LinkQ {
	return &linkQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type linkQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *linkQ) SaveURL(value storage.Link) (*storage.Link, error) {
	var result storage.Link
	stmt := sq.Insert(linkTableName).Columns("alias", "url").Values(value.Alias, value.URL).Suffix("returning id")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert url to db")
	}
	return &result, nil
}

func (q *linkQ) GetURL(alias string) (*storage.Link, error) {
	var result storage.Link
	err := q.db.Get(&result, q.sql.Select("*").From(linkTableName).Where("alias = ?", alias))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get url from db")
	}
	return &result, nil
}
