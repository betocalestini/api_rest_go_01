package controllers

import (
	"api_rest_go_01/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	//escrever na tela um array de bytes com a descrição Home Page
	w.Write([]byte("Home page"))
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	//encodar o retorno do w em json
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	//pegar o valor do ID com o mux
	vars := mux.Vars(r)
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		//comparar os valores, transformando ID em string
		if strconv.Itoa(personalidade.ID) == id {
			//retornar o jsonEncode de personalidade
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
