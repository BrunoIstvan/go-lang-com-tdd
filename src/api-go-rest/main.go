package main

import (
	"fmt"

	"github.com/BrunoIstvan/go-rest-api/database"
	"github.com/BrunoIstvan/go-rest-api/routes"
)

func main() {

	database.ConnectDatabase()

	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()

}
