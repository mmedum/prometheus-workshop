package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mmedum/prometheus-workshop/go-service/handlers/health"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func pong(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "pong"

	rand.Seed(time.Now().Unix())

	var responseCodes [3]int
	responseCodes[0] = 200
	responseCodes[1] = 500
	responseCodes[2] = 503

	responseCode := responseCodes[rand.Intn(len(responseCodes))]

	render.Status(r, responseCode)

	render.JSON(w, r, response)
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		middleware.RequestID,
	)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	router.Use(cors.Handler)

	router.Mount("/metrics", promhttp.Handler())
	router.Mount("/health", health.Routes())

	router.Route("/v1", func(r chi.Router) {
		r.Get("/ping", pong)
	})

	return router
}

func main() {
	router := Routes()
	port := 80

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
	log.Println("Running")
}
