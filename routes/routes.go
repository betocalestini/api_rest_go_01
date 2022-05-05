package routes

import (
	"api_rest_go_01/controllers"
	"log"
	"net/http"
	"os"
)

func HandleRequest() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}

	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
