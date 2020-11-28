package main

import (
	"fmt"
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/model"
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/resolvers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ArtemGretsov/golang-gqlgen-gorm-psql-example/graph/generated"
	_ "github.com/joho/godotenv"
	_ "gorm.io/driver/postgres"
	_ "gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	db := connectDatabase()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{DB: db}}))

	log.Printf("Server started! Port: " + port)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func connectDatabase() *gorm.DB {
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPort, dbPass)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error db connect")
	}

	mErr := db.AutoMigrate(
		&model.Day{},
		&model.Weather{},
		&model.WeatherDifference{},
		&model.Rate{},
		&model.RateDifference{},
	)

	if mErr != nil {
		log.Fatal("Migration Error")
	}

	return db
}
