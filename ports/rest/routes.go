package api

import (
	"log"
	"net/http"
	"time"

	"github.com/fmoral2/mux-api/application/model"

	"github.com/gorilla/mux"
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

func (a *EmpHandler) Routes(r *mux.Router) {
	appNew, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("mux-test"),
		newrelic.ConfigLicense("8e192759449a29e8ab34f9b9a3e4354e7ca1NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)

	r.HandleFunc(newrelic.WrapHandleFunc(appNew, "/api/employees", a.GetEmployees)).Methods("GET")
	r.Handle("/api/employees/filter", loggingMiddleware(http.HandlerFunc(a.GetFilterEmployees))).Methods("GET")
	r.Handle("/api/employees/{id}", loggingMiddleware(http.HandlerFunc(a.GetSingleEmployee))).Methods("GET")
	r.Handle("/api/employees/{id}", loggingMiddleware(http.HandlerFunc(a.DeleteEmployee))).Methods("DELETE")
	r.Handle("/api/employees/{id}", loggingMiddleware(http.HandlerFunc(a.UpdateEmployee))).Methods("PUT")
	r.Handle("/api/employees", loggingMiddleware(http.HandlerFunc(a.CreateEmployee))).Methods("POST")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s %v %v %v", r.Method, r.RequestURI, r.RemoteAddr, time.Since(startTime), r.Header.Get("User-Agent"), r.Body)
	})
}
