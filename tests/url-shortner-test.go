package tests

import (
	"net/http"
	"net/url"
	"path"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gavv/httpexpect/v2"
	"github.com/stretchr/testify/require"
	"github.com/url-shoter/iternal/http-server/handlers/url/save"
	"github.com/url-shoter/iternal/lib/api"
	"github.com/url-shoter/iternal/lib/random"
)

const (
	host = "localhost:8082"
)

func TestURLShortner(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	e := httpexpect.Default(t, u.String())

	e.POST("/url").
		WithJSON(save.Request{
			URL:   gofakeit.URL(),
			Alias: random.NewRandomString(10),
		}).
		WithBasicAuth("myuser", "mypassword").
		Expect().
		Status(200).
		JSON().
		Object().
		ContainsKey("alias")
}

func TestURLShortner_SaveRedirectRemove(t *testing.T) {
	testCases := []struct {
		name  string
		url   string
		alias string
		error string
	}{
		{
			name:  "Valid URL",
			url:   gofakeit.URL(),
			alias: gofakeit.Word() + gofakeit.Word(),
		},
		{
			name:  "Invalid URL",
			url:   "invalid_url",
			alias: gofakeit.Word(),
			error: "field URL is not a valid URL",
		},
		{
			name:  "Empty Alias",
			url:   gofakeit.URL(),
			alias: "",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			u := url.URL{
				Scheme: "http",
				Host:   host,
			}

			e := httpexpect.Default(t, u.String())

			req := e.POST("/url").
				WithJSON(save.Request{
					URL:   tc.url,
					Alias: tc.alias,
				}).
				WithBasicAuth("myuser", "mypassword").
				Expect().Status(http.StatusOK).
				JSON().Object()

			if tc.error != "" {
				req.NotContainsKey("alias")

				req.Value("error").String().IsEqual(tc.error)

				return
			}

			alias := tc.alias

			if tc.alias != "" {
				req.Value("alias").String().IsEqual(tc.alias)
			} else {
				req.Value("alias").String().NotEmpty()

				alias = req.Value("alias").String().Raw()
			}

			testRedirect(t, alias, tc.url)

			reqDel := e.DELETE("/"+path.Join("url", alias)).
				WithBasicAuth("myuser", "mypassword").
				Expect().Status(http.StatusOK).
				JSON().Object()

			reqDel.Value("status").String().IsEqual("ok")

			testRedirectNotFound(t, alias)
		})
	}
}

func testRedirect(t *testing.T, alias string, urlToRedirect string) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   alias,
	}

	redirectedToURL, err := api.GetRedirect(u.String())
	require.NoError(t, err)

	require.Equal(t, urlToRedirect, redirectedToURL)
}

func testRedirectNotFound(t *testing.T, alias string) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   alias,
	}

	_, err := api.GetRedirect(u.String())

	require.ErrorIs(t, err, api.ErrInvalidStatusCode)
}
