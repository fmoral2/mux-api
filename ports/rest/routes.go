package api

import (
	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/application/model"
	"github.com/newrelic/go-agent/v3/newrelic"
)

type EmpHandler struct {
	app EmpApp
}
type EmpApp interface {
	CreatingEmployeeDb(emp model.Employee) (string, error)
	GetEmployee(emp model.Employee) (model.Employee, error)
	GetFilterEmployees(emp model.Employee, pageParams model.PageRequest) (interface{}, error)
	GetEmployees(emp model.Employee) ([]model.Employee, error)
	DeleteEmployee(emp model.Employee) (model.Employee, error)
	UpdateEmployee(emp model.Employee) (model.Employee, error)
}

func MakeHandler(appEmp EmpApp) *EmpHandler {
	return &EmpHandler{app: appEmp}
}

func (a *EmpHandler) MakingRoutes(r *mux.Router) {
	appNew, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("mux-test"),
		newrelic.ConfigLicense("8e192759449a29e8ab34f9b9a3e4354e7ca1NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	r.HandleFunc(newrelic.WrapHandleFunc(appNew, "/api/employees", a.GetEmployees)).Methods("GET")
	//r.HandleFunc("/api/employees", a.GetEmployees).Methods("GET")
	r.HandleFunc("/api/employees/filter", a.GetFilterEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", a.GetSingleEmployee).Methods("GET")
	r.HandleFunc("/api/employees/{id}", a.DeleteEmployee).Methods("DELETE")
	r.HandleFunc("/api/employees/{id}", a.UpdateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees", a.CreateEmployee).Methods("POST")
}
