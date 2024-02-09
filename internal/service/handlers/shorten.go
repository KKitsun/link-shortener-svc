package handlers

import (
	"math/rand"
	"net/http"
	"time"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/KKitsun/link-shortener-svc/internal/service/requests"
)

type AliasResponse struct {
	Alias string `json:"alias"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) error
}

func Shorten(urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := requests.NewGetURL(r)
		if err != nil {
			ape.RenderErr(w, problems.BadRequest(err)...)
			return
		}

		alias := generateAlias()

		err = urlSaver.SaveURL(req.URL, alias)
		if err != nil {
			Log(r).WithError(err).Error("Error saving url to the database")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, AliasResponse{
			Alias: alias,
		})

	}
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
