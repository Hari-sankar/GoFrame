package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {

	//initilaze env files

	err := godotenv.Load(".env")

	if err != nil {
		log.Println(" Error : Error loading environment variables")

	}

}
