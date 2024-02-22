package handlers_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KKitsun/link-shortener-svc/internal/service/handlers"
	"github.com/KKitsun/link-shortener-svc/internal/storage"
	"github.com/KKitsun/link-shortener-svc/internal/storage/mocks"
	"github.com/go-chi/chi"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/distributed_lab/ape"
)

func TestGetFullHandler(t *testing.T) {
	cases := []struct {
		name       string
		alias      string
		url        string
		StatusCode int
		mockError  bool
	}{
		{
			name:       "Success1",
			alias:      "alias1",
			url:        "https://google.com",
			StatusCode: http.StatusOK,
		},
		{
			name:       "Success2",
			alias:      "alias2",
			url:        "https://youtu.be/G1IbRujko-A?si=BTtkfqEmNVckI-4S",
			StatusCode: http.StatusOK,
		},
		{
			name:       "Url Not Found",
			alias:      "alias4",
			url:        "",
			StatusCode: http.StatusInternalServerError,
			mockError:  true,
		},
	}

	for _, testCase := range cases {

		t.Run(testCase.name, func(t *testing.T) {

			mockLinkStorage := mocks.NewLinkStorage(t)
			mockLinkQ := mocks.NewLinkQ(t)

			mockLinkStorage.On("New").
				Return(mockLinkStorage).
				Once()

			mockLinkStorage.On("Link").
				Return(mockLinkQ).
				Once()

			if testCase.mockError {
				mockLinkQ.On("GetURL", testCase.alias).
					Return(&storage.Link{Alias: "", URL: ""}, nil).
					Once()
			} else {
				mockLinkQ.On("GetURL", testCase.alias).
					Return(&storage.Link{Alias: testCase.alias, URL: testCase.url}, nil).
					Once()
			}

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/%s", testCase.alias), nil)
			require.NoError(t, err)
			rr := httptest.NewRecorder()
			r := chi.NewRouter()
			r.Use(
				ape.CtxMiddleware(
					handlers.CtxDB(mockLinkStorage),
					handlers.CtxLog(mocks.NewFakeLogger().FakeLog()),
				),
			)
			r.Get("/{alias}", handlers.GetFull)

			r.ServeHTTP(rr, req)

			require.Equal(t, testCase.StatusCode, rr.Code)

			if testCase.StatusCode == http.StatusOK {
				body := rr.Body.String()
				var resp handlers.FullURLResponse
				require.NoError(t, json.Unmarshal([]byte(body), &resp))
				assert.Equal(t, testCase.url, resp.URL)
			}

		})
	}
}
