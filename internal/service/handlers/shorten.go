package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/KKitsun/link-shortener-svc/internal/service/requests"
	"github.com/KKitsun/link-shortener-svc/internal/storage"
)

type AliasResponse struct {
	Alias string `json:"alias"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	db := DB(r)

	req, err := requests.NewGetURL(r)
	if err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	alias := generateAlias()
	link := storage.Link{
		Alias: alias,
		URL:   req.URL,
	}

	_, err = db.Link().SaveURL(link)
	if err != nil {
		Log(r).WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, AliasResponse{
		Alias: alias,
	})

}

func generateAlias() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	alias := make([]byte, codeLength)
	for i := range alias {
		alias[i] = charset[rnd.Intn(len(charset))]
	}
	return string(alias)
}
