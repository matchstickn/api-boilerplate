package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/matchstickn/REPONAME/app/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	app := cmd.SetUpFiber()
	cmd.SetUpRouter(app)
	fmt.Println("Server is running on port 4000")
	app.Listen(":4000")
}
