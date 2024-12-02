package api

import (
	"goframe/pkg/database/postgres"
	"goframe/pkg/logger"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Run() {
	log.Println("starting server")

	err := godotenv.Load("./.env")
	if err != nil {
		log.Println(" Error : Error loading environment variables")

	}

	//initializing logger
	appLogger := logger.New()

	appLogger.Infof("Initialized logger")

	//Initialising db
	connectionURL := os.Getenv("DB_URL")

	psqlDB, err := postgres.NewPsqlDB(connectionURL)
	if err != nil {
		appLogger.Errorf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stat())
	}
	defer psqlDB.Close()

	// server.Initialise(os.Getenv("DB_HOST"), os.Getenv("DB_SCHEMA_USER"), os.Getenv("DB_SCHEMA_USER_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	// server.Run()

}
