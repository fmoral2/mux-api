package api

import (
	"github.com/gorilla/mux"
	application "github.com/morlfm/rest-api/application/employee"
)

type EmpHandler struct {
	app *application.App
}

func MakeHandler(appEmp *application.App) *EmpHandler {
	return &EmpHandler{app: appEmp}
}

func (a *EmpHandler) MakingRoutes(r *mux.Router) {
	// Initialize the router and endpoints

	r.HandleFunc("/api/employees", a.GetEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", a.GetSingleEmployee).Methods("GET")
	r.HandleFunc("/api/employees/{id}", a.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/api/employees/{id}", a.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees", a.CreateEmployee).Methods("POST")

}
