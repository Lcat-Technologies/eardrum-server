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
	"github.com/Lcat-Technologies/eardrum-postgres/merchant"
	"github.com/Lcat-Technologies/eardrum-postgres/postgresutils"
	"github.com/Lcat-Technologies/eardrum-postgres/transaction"
	"github.com/Lcat-Technologies/eardrum-postgres/user"
	"github.com/Lcat-Technologies/eardrum-postgres/device"
	"github.com/Lcat-Technologies/eardrum-server/auth"
	"github.com/Lcat-Technologies/eardrum-server/graph"
	"github.com/Lcat-Technologies/eardrum-server/phoneutils"
	"github.com/Lcat-Technologies/eardrum-server/pkg/jwt"
	"github.com/Lcat-Technologies/eardrum-server/shutdown"
	"github.com/go-chi/chi"
	"github.com/vektah/gqlparser/v2/gqlerror"

	//"github.com/joho/godotenv"
	customErrors "github.com/Lcat-Technologies/eardrum-interfaces/errors"
	"github.com/rs/cors"
	customLog "github.com/rs/zerolog/log"
)

var (
	postgresInstance postgresutils.PostgresInstance
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
		err.Extensions["code"] = publicErr.SystemCode
		err.Extensions["status"] = publicErr.HttpStatus
		err.Message = publicErr.Message
	}
	// ALWAYS log the FULL, wrapped error internally for debugging
	customLog.Error().Err(e).Msg("GraphQL Error Encountered")

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
	//Ensure environment variables are configured
	EnsureEnvVariables()

	go phoneutils.InitializeTwilio()
	go jwt.InitializeJwtSecretKey()

	//set IsShutdown to false
	s := false
	shutdown.IsShutdown = &s

	defaultPort := os.Getenv("DEFAULT_PORT")

	postgresInstance.Init(os.Getenv("POSTGRES_DBURL"))

	// Perform auto-migration for multiple models
	err := postgresInstance.Db.AutoMigrate(&user.User{}, &user.UnverifiedUser{}, &merchant.Merchant{}, &merchant.UnverifiedMerchant{}, &transaction.Transaction{}, &device.Device{})
	if err != nil {
		customLog.Fatal().Msg(fmt.Sprintf("Failed to auto-migrate database: %s", err))
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

	customLog.Info().Msg(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	customLog.Fatal().Msg(http.ListenAndServe(":"+port, router).Error())
}

// EnsureEnvVariables checks for required environment variables and exits if any are missing.
func EnsureEnvVariables() {
	requiredEnvs := []string{
		"POSTGRES_DBURL",
		"TRANSACTION_FEE_PERCENT",
		"JWT_SECRET_KEY",
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"TWILIO_VERIFY_SERVICE_SID",
		"DEFAULT_PORT",
	}

	var missingEnvs []string

	for _, env := range requiredEnvs {
		if val := os.Getenv(env); val == "" {
			missingEnvs = append(missingEnvs, env)
		}
	}

	// If there are any missing variables, log a single structured fatal event and crash
	if len(missingEnvs) > 0 {
		customLog.Fatal().
			Strs("missing_keys", missingEnvs).
			Msg("Application terminating due to missing required environment variables")
	}
}
