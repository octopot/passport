package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/passport/server/router"
)

func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1/tracker", func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/instruction", api.GetTrackerInstructionV1)
		r.Post("/fingerprint", api.PostTrackerFingerprintV1)
	})

	return r
}
