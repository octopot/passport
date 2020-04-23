package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"go.octolab.org/ecosystem/passport/internal/server/router"
)

// NewRouter returns configured `github.com/go-chi/chi` router.
func NewRouter(api router.Server) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Route("/api/v1/tracker", func(r chi.Router) {
		r.Use(cors.New(cors.Options{
			AllowOriginFunc:  func(origin string) bool { return true },
			AllowCredentials: true,
		}).Handler)
		r.Route("/instruction", func(r chi.Router) {
			r.Use(middleware.NoCache)
			r.Get("/", api.GetTrackerInstructionV1)
		})
		r.Route("/fingerprint", func(r chi.Router) {
			r.Post("/", api.PostTrackerFingerprintV1)
		})
	})

	return r
}
