package pkg

import (
	"log"
	"os"

	"github.com/bryanwsebaraj/mealswithfriends/pkg/controllers"
	"github.com/bryanwsebaraj/mealswithfriends/pkg/seed"
	"github.com/bryanwsebaraj/mealswithfriends/scripts"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Cannot load env. Through %v", err)
	}

	server.Initialize(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	go scripts.AddOnTime(server.DB)

	server.Run(":8080")

}
