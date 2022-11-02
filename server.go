package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/stephendatascientist/go-graphql-postgres-api/database"
	"github.com/stephendatascientist/go-graphql-postgres-api/entities"
	"github.com/stephendatascientist/go-graphql-postgres-api/graph"
	"github.com/stephendatascientist/go-graphql-postgres-api/graph/generated"
	"github.com/stephendatascientist/go-graphql-postgres-api/repositories"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.NewConnection(config)

	if err != nil {
		log.Fatal("could not load the database")
	}
	err = entities.MigrateEmployee(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	EmployeeRepository := repositories.EmployeeRepository{
		DB: db,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		EmployeeRepository: EmployeeRepository,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
