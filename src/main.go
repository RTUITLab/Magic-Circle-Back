package main

import (
	"github.com/0B1t322/Magic-Circle/server"
	log "github.com/sirupsen/logrus"
)

// @title Magic-Circle API
// @version 1.0
// @description This is a server to get projects from github
// @BasePath /api/magic-circle
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := server.StartServer(); err != nil {
		log.WithFields(
			log.Fields{
				"pkg": "main",
				"err": err,
			},
		).Panic("Failed to start server")
	}
}