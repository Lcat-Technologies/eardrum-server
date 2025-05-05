package graph

import (
	"github.com/GigaDesk/eardrum-graph/neo4jutils"
	"github.com/GigaDesk/eardrum-server/graph/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Sql *db.AutoGqlDB // this is the new line the package db is autogenerate by this plugin
	Neo4j *neo4jutils.Neo4jInstance
}
