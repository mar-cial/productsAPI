package app

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/mar-cial/productsApi/pkg/products"
)

// App is the repository for all app related things
type App struct {
	Products []products.Product
}

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplates, err := template.ParseGlob("templates/*.tmpl")
	if err != nil {
		log.Fatalln("could not parse templates")
	}

	parsedTemplates.ExecuteTemplate(w, "base", nil)
}

func (a *App) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	encoder.Encode(a.Products[:100])
}

func (a *App) GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	encoder := json.NewEncoder(w)
	id := vars["id"]

	for _, v := range a.Products {
		if v.StoreID == id {
			encoder.Encode(v)
		}
	}
}
