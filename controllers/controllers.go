package controllers

import (
	"api_rest_go_01/database"
	"api_rest_go_01/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//escrever na tela um array de bytes com a descrição Home Page
	w.Write([]byte("Home page"))
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	//criar a variavel lista de personalidades
	var p []models.Personalidade
	database.DB.Find(&p)
	//encodar o retorno do w em json
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	//pegar o valor do ID com o mux
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
	// for _, personalidade := range models.Personalidades {
	// 	//comparar os valores, transformando ID em string
	// 	if strconv.Itoa(personalidade.ID) == id {
	// 		//retornar o jsonEncode de personalidade
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }
}
