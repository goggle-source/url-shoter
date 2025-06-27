package delete

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/url-shoter/iternal/lib/api/response"
	"github.com/url-shoter/iternal/lib/logger/slo"
)

type DeleteURL interface {
	DeleteURL(alias string) error
}

func New(log *slog.Logger, delURL DeleteURL) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.Delete.New"

		log := slog.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, response.Error("invalid request"))

			return
		}

		err := delURL.DeleteURL(alias)
		if err != nil {
			log.Error("failed delete url", slo.Err(err))

			render.JSON(w, r, response.Error("internal error"))

			return
		}

		log.Info("delete url")

		render.JSON(w, r, response.StatusOK)
	}
}
