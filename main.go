package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avar348/golangrssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfic struct {
	DB *database.Queries
}

func main() {

	feed, err := fetchFeed("https://blog.boot.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(feed)
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the env")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DBURL is not found in the env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Problem connecting to the database", err)
	}
	db := database.New(conn)
	apiConfic := apiConfic{
		DB: database.New(conn),
	}

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiConfic.handleCreateUser)
	v1Router.Get("/users", apiConfic.middlewareAuth(apiConfic.handleGetUser))

	v1Router.Post("/feeds", apiConfic.middlewareAuth(apiConfic.handleCreateFeed))
	v1Router.Get("/feeds", apiConfic.getAllFeeds)
	v1Router.Post("/feed_follows", apiConfic.middlewareAuth(apiConfic.handleCreateFeedFollow))
	v1Router.Get("/feed_follows", apiConfic.middlewareAuth(apiConfic.handleGetFeedFollows))
	v1Router.Delete("/feed_follows/{id}", apiConfic.middlewareAuth(apiConfic.handleDeleteFeedFollows))
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server starting on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PORT:", portString)
}
