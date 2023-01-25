package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/adapters/rabbit"
	"github.com/morlfm/rest-api/adapters/repository"
	application "github.com/morlfm/rest-api/application/employee"
	api "github.com/morlfm/rest-api/ports/rest"
)

func main() {
	////cognito.Cognito()
	var ms runtime.MemStats
	printMemStat(ms)

	db := repository.CreateConnection()
	repository := repository.MakeRepository(db)

	app := application.MakeApplication(repository)
	rabbit.Publish(app)

	handler := api.MakeHandler(app)
	router := mux.NewRouter()
	handler.MakingRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func printMemStat(ms runtime.MemStats) {
	runtime.ReadMemStats(&ms)
	fmt.Println("--------------------------------------")
	fmt.Println("Memory Statistics Reporting time: ", time.Now())
	fmt.Println("--------------------------------------")
	fmt.Println("Bytes of allocated heap objects: ", ms.Alloc)
	fmt.Println("Total bytes of Heap object: ", ms.TotalAlloc)
	fmt.Println("Bytes of memory obtained from OS: ", ms.Sys)
	fmt.Println("Count of heap objects: ", ms.Mallocs)
	fmt.Println("Count of heap objects freed: ", ms.Frees)
	fmt.Println("Count of live heap objects", ms.Mallocs-ms.Frees)
	fmt.Println("Number of completed GC cycles: ", ms.NumGC)
	fmt.Println("--------------------------------------")
}
