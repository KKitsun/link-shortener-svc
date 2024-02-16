package postgres

import (
	"github.com/KKitsun/link-shortener-svc/internal/storage"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewLinkStorage(db *pgdb.DB) storage.LinkStorage {
	return &linkStorage{
		db: db.Clone(),
	}
}

type linkStorage struct {
	db *pgdb.DB
}

func (ls *linkStorage) New() storage.LinkStorage {
	return NewLinkStorage(ls.db)
}

func (ls *linkStorage) Link() storage.LinkQ {
	return newLinkQ(ls.db)
}
