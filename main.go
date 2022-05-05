package main

import (
	"api_rest_go_01/models"
	"api_rest_go_01/routes"
	env "api_rest_go_01/services"
	"fmt"
)

func main() {
	env.Load()
	models.Personalidades = []models.Personalidade{
		{ID: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{ID: 2, Nome: "Nome 1", Historia: "Historia 1"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
