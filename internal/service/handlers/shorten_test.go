package handlers_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/KKitsun/link-shortener-svc/internal/service/handlers"
	"github.com/KKitsun/link-shortener-svc/internal/storage"
	"github.com/KKitsun/link-shortener-svc/internal/storage/mocks"
	"github.com/go-chi/chi"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

func TestShortenHandler(t *testing.T) {
	cases := []struct {
		name       string
		url        string
		statusCode int
		mockError  bool
	}{
		{
			name:       "Success1",
			url:        "https://google.com",
			statusCode: http.StatusOK,
		},
		{
			name:       "Success2",
			url:        "https://youtu.be/G1IbRujko-A?si=BTtkfqEmNVckI-4S",
			statusCode: http.StatusOK,
		},
		{
			name:       "Empty URL",
			url:        "",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Invalid URL",
			url:        "some invalid URL",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Fail saving URL",
			url:        "https://google.com",
			statusCode: http.StatusInternalServerError,
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

			if testCase.statusCode == http.StatusOK {

				mockLinkStorage.On("Link").
					Return(mockLinkQ).
					Once()

				matchedBy := mock.MatchedBy(func(arg storage.Link) bool {
					return arg.URL == testCase.url
				})

				mockLinkQ.On("SaveURL", matchedBy).
					Return(&storage.Link{}, nil).
					Once()
			}

			if testCase.mockError {
				mockLinkStorage.On("Link").
					Return(mockLinkQ).
					Once()

				matchedBy := mock.MatchedBy(func(arg storage.Link) bool {
					return arg.URL == testCase.url
				})

				mockLinkQ.On("SaveURL", matchedBy).
					Return(&storage.Link{}, errors.Errorf("Some storage error")).
					Once()
			}

			input := fmt.Sprintf(`{"url":"%s"}`, testCase.url)

			req, err := http.NewRequest(http.MethodPost, "/", bytes.NewReader([]byte(input)))
			require.NoError(t, err)

			rr := httptest.NewRecorder()

			r := chi.NewRouter()
			r.Use(
				ape.CtxMiddleware(
					handlers.CtxDB(mockLinkStorage),
					handlers.CtxLog(mocks.NewFakeLogger().FakeLog()),
				),
			)
			r.Post("/", handlers.Shorten)

			r.ServeHTTP(rr, req)

			require.Equal(t, testCase.statusCode, rr.Code)

		})
	}
}
