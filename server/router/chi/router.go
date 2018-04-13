package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/kamilsk/passport/server/router"
	"github.com/rs/cors"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1/tracker", func(r chi.Router) {

		r.Use(cors.New(cors.Options{AllowCredentials: true}).Handler)

		r.Route("/instruction", func(r chi.Router) {
			r.Use(middleware.NoCache)
			r.Get("/", api.GetTrackerInstructionV1)
		})

		r.Route("/fingerprint", func(r chi.Router) {
			r.Use(middleware.AllowContentType("application/json"))
			r.Post("/", api.PostTrackerFingerprintV1)
		})
	})

	return r
}
