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
	devPort  = flag.Uint("dev-port", 9001, "port on which the dev server will listen")
	demoPort = flag.Uint("demo-port", 9002, "port on which the demo server will listen")
)

func main() {
	flag.Parse()
	demo, dev := chi.NewRouter(), chi.NewRouter()
	for _, r := range []*chi.Mux{demo, dev} {
		r.Use(middleware.Recoverer)
		r.Use(middleware.RealIP)
		r.Use(middleware.Logger)
	}
	{
		html := template.Must(template.New("cross-origin").ParseFiles("public/cross-origin.html"))
		demo.Get("/", func(rw http.ResponseWriter, req *http.Request) {
			rw.WriteHeader(http.StatusOK)
			html.Lookup("cross-origin.html").
				Execute(rw, struct{ DevHost string }{DevHost: fmt.Sprintf("localhost:%d", *devPort)})
		})
		demo.Get("/*", func(rw http.ResponseWriter, req *http.Request) { rw.WriteHeader(http.StatusOK) })
		go func() {
			log.Printf("starting demo server at :%d, use `open http://localhost:%[1]d/` to view it\n", *demoPort)
			http.ListenAndServe(fmt.Sprintf(":%d", *demoPort), demo)
		}()
	}
	{
		js := template.Must(template.New("cross-origin").ParseFiles("public/js/passport.js"))
		dev.Get("/js/*", func(rw http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/js/passport.js" {
				rw.WriteHeader(http.StatusOK)
				js.Lookup("passport.js").
					Execute(rw, struct {
						// issue #19
						EncryptedMarker string

						Endpoint string
					}{
						EncryptedMarker: "demo-cross-origin",

						Endpoint: fmt.Sprintf("http://localhost:%d/api/v1/tracker/fingerprint", *devPort),
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
			http.SetCookie(rw, &http.Cookie{Name: "marker", Value: "demo-cross-origin", Path: "/",
				HttpOnly: true, Secure: true})
			io.Copy(rw, f)
		})
		dev.Route("/api/v1/tracker/fingerprint", func(r chi.Router) {
			r.Use(cors.New(cors.Options{AllowCredentials: true, Debug: true}).Handler)
			r.Post("/", func(rw http.ResponseWriter, req *http.Request) {
				log.Println(req.Cookie("marker"))
				time.Sleep(250 * time.Millisecond)
				io.Copy(ioutil.Discard, req.Body)
				req.Body.Close()
			})
		})
	}
	log.Printf("starting dev server at :%d\n", *devPort)
	http.ListenAndServe(fmt.Sprintf(":%d", *devPort), dev)
}
