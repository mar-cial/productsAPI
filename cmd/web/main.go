package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/mar-cial/productsApi/pkg/app"
	"github.com/mar-cial/productsApi/pkg/handlers"
	"github.com/mar-cial/productsApi/pkg/products"
)

const (
	port = ":8080"
)

func RunServer(r *mux.Router) {
	// Add your routes as needed

	srv := &http.Server{
		Addr: port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

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

	srv := http.Server{
		Addr:         port,
		Handler:      router,
		WriteTimeout: time.Second * 12,
		ReadTimeout:  time.Second * 12,
		IdleTimeout:  time.Second * 60,
	}

	fmt.Printf("running app on port %s\n", port)
	log.Fatalln(srv.ListenAndServe())
}
