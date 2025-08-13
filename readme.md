# API-Boiler

Uses Fiber, Swagger, Testify, and Godotenv and that's it. Extend it however you like.


## Installation

Just copy this repo and initalized a go module

```
  go mod init REPONAME
  go mod tidy
```
    
## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`ENABLE_IP_VALIDATION`

`APP_NAME`

`PORT`

**OPTIONAL**

`RATE_LIMITING`

## Running the Thing

Just use the standard go run command in the root directory

```
  go run main.go
```


## Extras

 - You can delete the api/index.go file because that's for vercel deployments only
 - When using swagger generate using 'swag init -g app/controller/controller.go -o ./app/docs
'


## Acknowledgements

 - github.com/gofiber/fiber/v2
 - github.com/joho/godotenv
 - github.com/stretchr/testify
 - github.com/swaggo/swag
