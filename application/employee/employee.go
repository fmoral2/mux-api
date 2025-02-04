package application

import (
	"encoding/json"
	"net/http"

	"github.com/fmoral2/mux-api/adapters/repository"
	"github.com/fmoral2/mux-api/application/model"
)

type App struct {
	empRepository *repository.EmployeePG
}

func MakeApplication(empDb *repository.EmployeePG) *App {
	return &App{empRepository: empDb}
}

func (a *App) CreatingEmployeeDb(emp model.Employee) (string, error) {
	id, err := a.empRepository.CreateEmployeeDb(emp.Name, emp.Location, emp.Wage, emp.Role)
	return id, err
}

func (a *App) GetEmployee(emp model.Employee) (model.Employee, error) {
	emp, err := a.empRepository.GetEmployeeDb(emp.ID)
	return emp, err
}

func (a *App) GetEmployees(emp model.Employee) ([]model.Employee, error) {
	emps, err := a.empRepository.GetEmployeesDb()
	return emps, err
}

func (a *App) GetFilterEmployees(emp model.Employee, pageParams model.PageRequest) (interface{}, error) {
	pageParam := pageParams.WithDefaultValues()
	emps, err := a.empRepository.GetEmployeesFilterDb(pageParam.Page, pageParam.Size)
	return emps, err
}

func (a *App) DeleteEmployee(emp model.Employee) (model.Employee, error) {
	emp, err := a.empRepository.DeleteEmployeeDb(emp.ID)
	return emp, err
}

func (a *App) UpdateEmployee(emp model.Employee) (model.Employee, error) {
	empUp, err := a.empRepository.UpdateEmployeeDb(emp)
	return empUp, err
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		return
	}
	w.WriteHeader(code)

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}
