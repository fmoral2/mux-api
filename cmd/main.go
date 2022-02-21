package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	application "github.com/morlfm/rest-api/application/employee"
	"github.com/morlfm/rest-api/application/repository"
	"github.com/morlfm/rest-api/ports/api"
	"github.com/morlfm/rest-api/rabbit"
)

func main() {

	// rabbit.Consumer()
	db := repository.CreateConnection()
	repository := repository.MakeRepository(db)

	app := application.MakeApplication(repository)
	rabbit.Publish(app)

	handler := api.MakeHandler(app)

	router := mux.NewRouter()
	handler.MakingRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
