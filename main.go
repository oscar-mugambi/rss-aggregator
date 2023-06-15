package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in env")
	}
	// creates a new router object
	router := chi.NewRouter()

	// add a cors config to our router
	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	// we need to hook up our router
	// using the chi router we hook up the handler to a specific HTTP method and path

	v1Router := chi.NewRouter()
	// connect the handlerReadiness function to the /readyPath
	// v1Router.HandleFunc("/heatlhz", handlerReadiness)
	v1Router.Get("/heatlhz", handlerReadiness)

	v1Router.Get("/err", handlerErr)

	// we nest a v1 router under the /v1 path
	router.Mount("/v1", v1Router)

	// connect router to a HTTP server
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	// ListenAndServe will block. Code stops here and starts handling requests
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
