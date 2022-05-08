package main

import (
	"log"

	"github.com/mar-cial/productsApi/pkg/app"
	"github.com/mar-cial/productsApi/pkg/handlers"
	"github.com/mar-cial/productsApi/pkg/products"
	"github.com/mar-cial/productsApi/pkg/server"
)

func main() {

	products, err := products.LoadProducts()
	if err != nil {
		log.Fatalln(err)
	}

	app := app.App{
		Products: products,
	}

	router, err := handlers.AddRouter(&app)
	if err != nil {
		log.Fatalln("could not add router: ", err)
	}

	server.RunServer(router)
}
