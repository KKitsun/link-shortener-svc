package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

type FullURLResponse struct {
	URL string `json:"url"`
}

type URLGetter interface {
	GetURL(alias string) ([]string, error)
}

func GetFull(urlGetter URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			Log(r).Error("Error getting alias from get request")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		urlFromDB, err := urlGetter.GetURL(alias)
		if err != nil {
			Log(r).WithError(err).Error("Error getting url from db")
			ape.RenderErr(w, problems.InternalError())
			return
		}
		if urlFromDB == nil {
			Log(r).WithError(err).Error("Error getting url from db URL IS NIL")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, FullURLResponse{
			URL: urlFromDB[0],
		})

	}
}
