package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey   ctxKey = iota
	urlAliasKey string = ""
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxUrlAlias() func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, urlAliasKey, "alias")
	}
}

func UrlAlias(r *http.Request) string {
	return r.Context().Value(urlAliasKey).(string)
}
