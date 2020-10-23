package main

import (
	"log"
	"recipe-api/routers"
	"recipe-api/utils"
)

func main() {
	router := routers.GetRouter()
	port := utils.EnvVar("SERVER_PORT")
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
