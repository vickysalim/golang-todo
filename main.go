package main

import (
	"todo-app/config"
	"todo-app/routes"
)

func main() {
    config.ConnectDatabase()
    r := routes.SetupRouter()
    r.Run(":8080")
}
