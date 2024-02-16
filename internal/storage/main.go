package storage

type LinkStorage interface {
	New() LinkStorage

	Link() LinkQ
}
