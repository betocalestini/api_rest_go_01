package main

import (
	"api_rest_go_01/models"
	"api_rest_go_01/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Nome 1", Historia: "Historia 1"},
		{Nome: "Nome 1", Historia: "Historia 1"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
