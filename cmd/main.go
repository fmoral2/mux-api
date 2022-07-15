package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/adapters/rabbit"
	"github.com/morlfm/rest-api/adapters/repository"
	application "github.com/morlfm/rest-api/application/employee"
	api "github.com/morlfm/rest-api/ports/rest"
)

func main() {

	db := repository.CreateConnection()
	repository := repository.MakeRepository(db)

	app := application.MakeApplication(repository)
	rabbit.Publish(app)

	handler := api.MakeHandler(app)
	router := mux.NewRouter()
	handler.MakingRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
