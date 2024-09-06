package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Shokolat1/rssAggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	// Get the port and db_url env variables and store it
	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not found in the .env file")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the .env file")
	}

	// Connection to DB
	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to DB:", err)
	}
	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	// Create router
	router := chi.NewRouter()

	// Use CORS in router config
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Connect the handleReadiness function (sending an ok response when receiving petition to "/ready")
	v1Router := chi.NewRouter()

	// Handle for every HTTP verb
	// v1Router.HandleFunc("/health", handlerReadiness)

	// Handle for GET requests
	v1Router.Get("/health", handlerReadiness)

	// Handle errors (not a valid route)
	v1Router.Get("/err", handlerError)

	// Create User
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	// Get user based on api key
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	// Create Feed
	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	// Get all Feeds
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	// Mount all the routes inside v1Router into the general router. An example of how routes will look like is: localhost:8000/v1/health
	router.Mount("/v1", v1Router)

	// Create server (routing handler and address to listen to)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	// Start listening to port calls, and hand out error if there is one
	log.Printf("Server starting on port: %v", portStr)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT: ", portStr)

}
