package product_api

import (
	"Microservice/product-api/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	env.Parsel()

	l := log.New(os.Stdout, "products-api ", log.LstdFlags)

	//creating the handlers
	ph := handlers.NewProducts(l)

	//creating a new servmux and registering the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	//creating a new server
	s := http.Server{
		Addr:         *bindAddress,
		Handler:      sm,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
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
