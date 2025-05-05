package postgreutils

import (
	"fmt"

	"github.com/GigaDesk/eardrum-server/graph/db"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresInstance struct {
	Dborm db.AutoGqlDB
}

func (p *PostgresInstance) Init(url string) {
	log.Trace().Str("url", url).Msg("initializing postgresql")
	dbInstance, dbError := gorm.Open(postgres.Open(url), &gorm.Config{})
	if dbError != nil {
		log.Fatal().Str("url", url).Msg(fmt.Sprintf("problem initializing postgresql database: %s", dbError))
	}
	dborm := db.NewAutoGqlDB(dbInstance)
	dborm.Init()
	p.Dborm = dborm
	log.Trace().Str("url", url).Msg("completed postgresql initialization")
}
