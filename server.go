package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-postgres/postgresutils"
	"github.com/GigaDesk/eardrum-server/auth"
	"github.com/GigaDesk/eardrum-server/graph"
	"github.com/GigaDesk/eardrum-server/phoneutils"
	"github.com/GigaDesk/eardrum-server/pkg/jwt"
	"github.com/GigaDesk/eardrum-server/shutdown"
	"github.com/go-chi/chi"

	//"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
)

var (
	postgresInstance postgresutils.PostgresInstance
	neo4jInstance  neo4jutils.Neo4jInstance
)


func main() {

	// Find .env file
	
	/*err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error loading .env file: %s", err))
	}*/

	go phoneutils.InitializeTwilio()
	go jwt.InitializeJwtSecretKey()

	//set IsShutdown to false
    s:= false
	shutdown.IsShutdown = &s

	defaultPort := os.Getenv("DEFAULT_PORT")

	postgresInstance.Init(os.Getenv("POSTGRES_DBURL"))
	neo4jerr:=neo4jInstance.Init(os.Getenv("NEO4J_DBURI"), os.Getenv("NEO4J_DBUSER"), os.Getenv("NEO4J_DBPASSWORD"))

	if neo4jerr !=nil {
		log.Fatal().Msg(fmt.Sprintf("problem initializing neo4j database: %s", neo4jerr))
	}
	
	defer neo4jInstance.Driver.Close(neo4jInstance.Ctx)

	port := defaultPort
	router := chi.NewRouter()
	
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}, // Include "Authorization"
	})
    
	router.Use(c.Handler)
	router.Use(auth.Middleware())

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Sql: &postgresInstance, Neo4j: &neo4jInstance }}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", server)

	log.Info().Msg(fmt.Sprintf("connect to http://localhost:%s/ for GraphQL playground", port))
	log.Fatal().Msg(http.ListenAndServe(":"+port, router).Error())
}
