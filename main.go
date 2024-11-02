package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	// Import the handler_readiness function
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/ready", HandlerReadiness)
	v1Router.Get("/error", HandlerError)

	router.Mount("/v1", v1Router)

	// srv := &http.Server{
	// 	Handler: router,
	// 	Addr:    ":" + port,
	// }
	log.Printf("Server is running on port %v", port)
	// err := srv.ListenAndServe()
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// }
	http.ListenAndServe(":"+port, router)

}
