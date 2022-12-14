package main

import (
	"Microservice/product-api/handlers"
	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind Address for the server")

func main() {
	env.Parse()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	//creating the handlers
	ph := handlers.NewProducts(l)

	//creating a new server mux and registering the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	//creating a new server
	s := http.Server{
		Addr:         *bindAddress,      //configuring the bind address
		Handler:      sm,                //setting the default handler
		ErrorLog:     l,                 //setting the logger for the server
		ReadTimeout:  5 * time.Second,   //maximum time to read request from the client
		WriteTimeout: 10 * time.Second,  //maximum time to write response for the client
		IdleTimeout:  120 * time.Second, //maximum time for connections using TCP Keep-Alive
	}

	// starting the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting the server")
			os.Exit(1)
		}
	}()

	//traping sigtermor interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)
}
