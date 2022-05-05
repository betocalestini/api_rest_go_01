package main

import (
	"api_rest_go_01/database"
	"api_rest_go_01/routes"
	env "api_rest_go_01/services"
	"fmt"
)

func main() {
	env.Load()

	database.ConectaBandoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
