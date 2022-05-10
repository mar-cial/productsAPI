package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mar-cial/productsApi/pkg/app"
)

func AddRouter(app *app.App) (*mux.Router, error) {
	mux := mux.NewRouter()
	mux.HandleFunc("/", app.Home).Methods(http.MethodGet)
	mux.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	api := mux.PathPrefix("/api").Subrouter().StrictSlash(true)
	api.HandleFunc("/products", app.GetAllProducts).Methods(http.MethodGet)
	api.HandleFunc("/products/{id:[0-9]+}", app.GetSingleProduct).Methods(http.MethodGet)
	return mux, nil
}
