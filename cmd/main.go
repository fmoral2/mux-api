package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/trace"
	"time"

	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/adapters/rabbit"
	"github.com/morlfm/rest-api/adapters/repository"
	application "github.com/morlfm/rest-api/application/employee"
	api "github.com/morlfm/rest-api/ports/rest"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("Failed to create trace file: %v", err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("Failed to close trace file: %v", err)
		}
	}()

	err = trace.Start(f)
	if err != nil {
		log.Fatalf("Failed to start tracing: %v", err)
	}
	defer trace.Stop()

	// //cognito.Cognito()
	go func() {
		for {
			var ms runtime.MemStats
			printMemStat(ms)
			time.Sleep(3 * time.Minute)
		}
	}()

	db := repository.CreateConnection()
	rep := repository.MakeRepository(db)

	app := application.MakeApplication(rep)
	rabbit.Publish(app)

	handler := api.MakeHandler(app)
	router := mux.NewRouter()
	handler.Routes(router)
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
