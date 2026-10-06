package graph

import (
	"github.com/Lcat-Technologies/eardrum-postgres/postgresutils"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Sql   *postgresutils.PostgresInstance
}
