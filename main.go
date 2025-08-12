package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matchstickn/REPONAME/app/cmd"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := cmd.SetUpFiber()
	cmd.SetUpRouter(app)
	fmt.Println("Server is running on port " + os.Getenv("PORT"))
	app.Listen(":" + os.Getenv("PORT"))
}
