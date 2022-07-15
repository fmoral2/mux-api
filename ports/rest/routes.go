package api

import (
	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/application/model"
)

type EmpHandler struct {
	app EmpApp
}
type EmpApp interface {
	CreatingEmployeeDb(emp model.Employee) (string, error)
	GetEmployee(emp model.Employee) (model.Employee, error)
	GetEmployees(emp model.Employee) ([]model.Employee, error)
	DeleteEmployee(emp model.Employee) (model.Employee, error)
	UpdateEmployee(emp model.Employee) (model.Employee, error)
}

func MakeHandler(appEmp EmpApp) *EmpHandler {
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
