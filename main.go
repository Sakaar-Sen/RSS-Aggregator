package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Sakaar-Sen/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// feed, err := urlToFeed("https://wagslane.dev/index.xml")
	// if err != nil {
	// 	log.Fatalf("Error fetching feed: %v", err)
	// }
	// fmt.Println(feed)

	godotenv.Load()
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DB_URL")

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	apiCfg := &apiConfig{
		DB: database.New(conn),
	}

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
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.HandlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.HandlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.HandlerDeleteFeedFollows))

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.HandlerGetPostsForUser))

	router.Mount("/v1", v1Router)

	go startScraping(apiCfg.DB, 10, 1*time.Minute)
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
