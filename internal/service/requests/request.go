package requests

import (
	"encoding/json"
	"net/http"

	. "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type URLRequest struct {
	URL string `json:"url" validate:"required,url"`
}

func NewGetURL(r *http.Request) (URLRequest, error) {
	var request URLRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, validateGetURL(request)
}

func validateGetURL(request URLRequest) error {
	data := &request

	return ValidateStruct(data,
		Field(&data.URL, Required),
		Field(&data.URL, is.URL),
	)
}
