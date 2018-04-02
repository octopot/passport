// +build demo-cross-origin

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

var (
	apiHost  = flag.String("api-host", "passport.local", "host name of the api server")
	apiPort  = flag.Uint("api-port", 8080, "port on which the api server will listen")
	siteHost = flag.String("site-host", "demo.local", "host name of the demo site")
	sitePort = flag.Uint("site-port", 80, "port on which the demo site will listen")
)

func main() {
	flag.Parse()
	site, api := chi.NewRouter(), chi.NewRouter()
	for _, r := range []*chi.Mux{site, api} {
		r.Use(middleware.Recoverer)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
	}
	{
		html := template.Must(template.New("cross-origin").ParseFiles("env/cross-origin/index.html"))
		site.Get("/", func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			html.Lookup("index.html").
				Execute(rw, struct{ APIHost string }{APIHost: *apiHost})
		})
		go func() {
			log.Printf("starting demo site at :%d, use `open http://%s/` to view it\n", *sitePort, *siteHost)
			http.ListenAndServe(fmt.Sprintf(":%d", *sitePort), site)
		}()
	}
	{
		js := template.Must(template.New("cross-origin").ParseFiles("public/js/passport.js"))
		api.Get("/js/*", func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "application/javascript")

			if req.URL.Path == "/js/passport.js" {
				log.Println("- set cookie")
				http.SetCookie(rw, &http.Cookie{Name: "marker", Value: "cross-origin", Path: "/", HttpOnly: true})
				log.Println("marker=cross-origin <nil>")
				log.Println("-")

				rw.WriteHeader(http.StatusOK)
				js.Lookup("passport.js").
					Execute(rw, struct{ Endpoint string }{
						Endpoint: fmt.Sprintf("http://%s/api/v1/tracker/fingerprint", *apiHost),
					})
				return
			}

			f, err := os.Open("public" + req.URL.Path)
			if err != nil {
				log.Println("unknown resource", req.URL.Path)
				rw.WriteHeader(http.StatusAccepted)
				return
			}
			defer f.Close()
			io.Copy(rw, f)
		})
		api.Route("/api/v1/tracker/fingerprint", func(r chi.Router) {
			r.Use(cors.New(cors.Options{AllowCredentials: true, Debug: true}).Handler)
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) {
				log.Println("- get cookie")
				log.Println(req.Cookie("marker"))
				log.Println("-")

				time.Sleep(250 * time.Millisecond)
				io.Copy(ioutil.Discard, req.Body)
				req.Body.Close()
			})
		})
	}
	log.Printf("starting api server at :%d\n", *apiPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *apiPort), api)
}
