package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-postgres/merchant"
	"github.com/GigaDesk/eardrum-postgres/postgresutils"
	"github.com/GigaDesk/eardrum-postgres/transaction"
	"github.com/GigaDesk/eardrum-postgres/user"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/graph"
	"github.com/GigaDesk/eardrum-server/phoneutils"
	"github.com/GigaDesk/eardrum-server/pkg/jwt"
	"github.com/GigaDesk/eardrum-server/shutdown"
	"github.com/go-chi/chi"
	"github.com/vektah/gqlparser/v2/gqlerror"

	//"github.com/joho/godotenv"
	customErrors "github.com/GigaDesk/eardrum-interfaces/errors"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

var (
	postgresInstance postgresutils.PostgresInstance
	neo4jInstance    neo4jutils.Neo4jInstance
)

func CustomErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := graphql.DefaultErrorPresenter(ctx, e)

	// Attempt to unwrap the error to your custom PublicError struct
	var publicErr *customErrors.PublicError
	if errors.As(e, &publicErr) {
		if err.Extensions == nil {
			err.Extensions = make(map[string]interface{})
		}
		// Copy the structured fields to the GraphQL extensions
		err.Extensions["code"] = publicErr.Code
		err.Extensions["reason"] = publicErr.Reason
		err.Message = publicErr.Message
	}
    // ALWAYS log the FULL, wrapped error internally for debugging
    log.Error().Err(e).Msg("GraphQL Error Encountered")

	return err
}

func main() {

	//Find .env file
    /*
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error loading .env file: %s", err))
	}
 */
	go phoneutils.InitializeTwilio()
	go jwt.InitializeJwtSecretKey()

	//set IsShutdown to false
	s := false
	shutdown.IsShutdown = &s

	defaultPort := os.Getenv("DEFAULT_PORT")

	postgresInstance.Init(os.Getenv("POSTGRES_DBURL"))

	// Perform auto-migration for multiple models
	err := postgresInstance.Db.AutoMigrate(&user.User{}, &user.UnverifiedUser{}, &merchant.Merchant{}, &merchant.UnverifiedMerchant{}, &product.Product{}, &product.Category{}, &transaction.Transaction{}, &transaction.Purchase{})
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to auto-migrate database: %s", err))
	}

	port := defaultPort
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}, // Include "Authorization"
	})

	router.Use(c.Handler)
	router.Use(auth.Middleware())

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &postgresInstance}}))
    server.SetErrorPresenter(CustomErrorPresenter)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Info().Msg(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	log.Fatal().Msg(http.ListenAndServe(":"+port, router).Error())
}
