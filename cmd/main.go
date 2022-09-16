package main

import (
	"log"
	"net/http"

	"github.com/newrelic/go-agent/v3/newrelic"

	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/adapters/rabbit"
	"github.com/morlfm/rest-api/adapters/repository"
	application "github.com/morlfm/rest-api/application/employee"
	api "github.com/morlfm/rest-api/ports/rest"
)

func main() {

	_, _ = newrelic.NewApplication(
		newrelic.ConfigAppName("TEST-mux"),
		newrelic.ConfigLicense("8e192759449a29e8ab34f9b9a3e4354e7ca1NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	db := repository.CreateConnection()
	repository := repository.MakeRepository(db)

	app := application.MakeApplication(repository)
	rabbit.Publish(app)

	handler := api.MakeHandler(app)
	router := mux.NewRouter()
	handler.MakingRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}
