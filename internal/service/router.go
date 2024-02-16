package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/KKitsun/link-shortener-svc/internal/config"
	"github.com/KKitsun/link-shortener-svc/internal/service/handlers"
	"github.com/KKitsun/link-shortener-svc/internal/storage/postgres"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(postgres.NewLinkStorage(cfg.DB())),
		),
	)
	r.Route("/integrations/link-shortener-svc", func(r chi.Router) {
		r.Post("/", handlers.Shorten)
		r.Get("/{alias}", handlers.GetFull)
	})

	return r
}
