package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/KKitsun/link-shortener-svc/internal/service/requests"
)

type FullURLResponse struct {
	URL string `json:"url"`
}

func GetFull(w http.ResponseWriter, r *http.Request) {
	db := DB(r)

	alias, err := requests.NewGetAlias(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	link, err := db.Link().GetURL(alias)
	if err != nil {
		Log(r).WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	if link.URL == "" {
		Log(r).Error("url not found")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, FullURLResponse{
		URL: link.URL,
	})

}
