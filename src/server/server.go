package server

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/0B1t322/Magic-Circle/config"
	"github.com/0B1t322/Magic-Circle/controllers/auth"
	"github.com/0B1t322/Magic-Circle/controllers/direction"
	"github.com/0B1t322/Magic-Circle/controllers/institute"
	"github.com/0B1t322/Magic-Circle/controllers/profile"
	"github.com/0B1t322/Magic-Circle/controllers/root"
	"github.com/0B1t322/Magic-Circle/controllers/sector"
	"github.com/0B1t322/Magic-Circle/db"

	log "github.com/sirupsen/logrus"
)

type Controllers struct {
	Sector    *sector.SectorController
	Profile   *profile.ProfileController
	Institute *institute.InstituteController
	Direction *direction.DirectionController
	Root      *root.RootController
	Auth      *auth.AuthController
}

func StartServer() error {
	config := config.GetConfig()

	client, err := db.CreateClient(config.DB.DBURI)
	if err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err":  err,
			},
		).Error("Failed to crete db client")

		return err
	}

	if err := client.Schema.Create(
		context.Background(),
		schema.WithForeignKeys(true),
		schema.WithDropColumn(true),
	); err != nil {
		log.WithFields(
			log.Fields{
				"func": "StartServer",
				"err":  err,
			},
		).Error("Failed to create schema")

		return err
	}

	controllers := &Controllers{
		Sector:    sector.New(client),
		Profile:   profile.New(client),
		Direction: direction.New(client),
		Institute: institute.New(client),
		Root:      root.New(client),
		Auth: auth.New(
			client,
			config.App.AccessSecret,
			config.App.RefreshSecret,
			config.App.SuperAdminLogin,
			config.App.SuperAdminPasswd,
		),
	}

	r := NewRouter(controllers)
	log.Infof("Starting server on %s", config.App.Port)
	return r.Run(fmt.Sprintf(":%s", config.App.Port))
}
