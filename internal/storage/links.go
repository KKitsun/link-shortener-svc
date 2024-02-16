package storage

type LinkQ interface {
	SaveURL(value Link) (*Link, error)
	GetURL(string) (*Link, error)
}

type Link struct {
	ID    int64  `db:"id"`
	Alias string `db:"alias"`
	URL   string `db:"url"`
}
