package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	App *App
	DB  *DB
}

type App struct {
	Port             string `envconfig:"MAGICCIRCLE_APP_PORT" default:"8080"`
	SuperAdminLogin  string `envconfig:"MAGICCIRCLE_APP_SUPERADMIN_LOGIN" default:"root"`
	SuperAdminPasswd string `envconfig:"MAGICCIRCLE_APP_SUPERADMIN_PASSWORD" default:"root"`
	AccessSecret     string `envconfig:"MAGICCIRCLE_APP_ACCESS_SECRET" default:"access_secret"`
	RefreshSecret    string `envconfig:"MAGICCIRCLE_APP_REFRESH_SECRET" default:"refresh_secret"`
}

type DB struct {
	DBURI string `envconfig:"MAGICCIRCLE_DB_URI" default:"root:root@/magic_circle"`
}

func GetConfig() *Config {
	var config Config
	{
		if err := godotenv.Load("./.env"); err != nil {
			log.Warn("Don't find env file")
		}

		if err := envconfig.Process("magiccircle", &config); err != nil {
			log.WithFields(
				log.Fields{
					"function": "envconfig.Process",
					"error":    err,
				},
			).Fatal("Can't read env vars, shutting down...")
		}
	}

	return &config
}
