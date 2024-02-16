package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/KKitsun/link-shortener-svc/internal/storage"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDB(entry storage.LinkStorage) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCtxKey, entry)
	}
}

func DB(r *http.Request) storage.LinkStorage {
	return r.Context().Value(dbCtxKey).(storage.LinkStorage).New()
}
