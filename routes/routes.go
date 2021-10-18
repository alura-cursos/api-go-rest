package routes

import (
	"log"
	"net/http"

	"github.com/guilhermeonrails/go-rest-api/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
