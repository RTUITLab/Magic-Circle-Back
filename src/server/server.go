package server

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/0B1t322/Magic-Circle/config"
	"github.com/0B1t322/Magic-Circle/controllers/sector"
	"github.com/0B1t322/Magic-Circle/db"

	log "github.com/sirupsen/logrus"
)

type Controllers struct {
	Sector *sector.SectorController
}

func StartServer() error {
	config := config.GetConfig()

	client, err := db.CreateClient(config.DB.DBURI)
	if err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err": err,
			},
		).Error("Failed to crete db client")

		return err
	}

	if err := client.Schema.Create(
		context.Background(),
		schema.WithForeignKeys(true),
	); err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err": err,
			},
		).Error("Failed to create schema")

		return err
	}

	controllers := &Controllers{
		Sector: sector.New(client),
	}


	r := NewRouter(controllers)
	return r.Run(fmt.Sprintf(":%s", config.App.Port))
}