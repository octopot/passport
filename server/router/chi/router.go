package chi

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(api interface{}) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	notImplemented := func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusNotImplemented) }

	r.Route("/api/v1/tracker", func(r chi.Router) {
		r.Use(middleware.NoCache)

		r.Get("/instruction", notImplemented)
		r.Post("/fingerprint", notImplemented)
	})

	return r
}
