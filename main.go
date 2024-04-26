package main

import (
	"filesProcessor/router"
	"log"
)

func main() {
	mainRouter := router.NewRouter()
	if err := mainRouter.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
