package health

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", GetHealth)
	return router
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]string)
	res["health"] = "OK"
	render.JSON(w, r, res)
}
