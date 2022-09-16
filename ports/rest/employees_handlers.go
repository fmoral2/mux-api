package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/morlfm/rest-api/adapters/rabbit"
	application "github.com/morlfm/rest-api/application/employee"
	"github.com/morlfm/rest-api/application/model"
)

var (
	emp model.Employee
)

func (a *EmpHandler) GetSingleEmployee(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	e := model.Employee{ID: id}

	if _, err := a.app.GetEmployee(e); err != nil {
		switch err {
		case sql.ErrNoRows:
			application.RespondWithError(w, http.StatusBadRequest, "not found")
			return
		default:
			application.RespondWithError(w, http.StatusBadRequest, "wrong id")
			return
		}
	}
	emp, _ := a.app.GetEmployee(e)

	json.NewEncoder(w).Encode(emp)
}

func (a *EmpHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	e := model.Employee{}
	emps, err := a.app.GetEmployees(e)

	if err != nil {

		return
	}
	json.NewEncoder(w).Encode(emps)
}

func (a *EmpHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	//set header
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// validating id format
	id := params["id"]
	if _, err := uuid.Parse(id); err != nil {
		application.RespondWithError(w, http.StatusBadRequest, "wrong id")
		return
	}

	// validating id in Db
	e := model.Employee{ID: id}
	if _, err := a.app.DeleteEmployee(e); err != nil {
		application.RespondWithError(w, http.StatusInternalServerError, "not found")
		return
	}
	application.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *EmpHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// set header
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	if _, err := uuid.Parse(id); err != nil {
		application.RespondWithError(w, http.StatusBadRequest, "wrong id")
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&emp); err != nil {
		application.RespondWithError(w, http.StatusBadRequest, "Invalid resquest")
	}
	defer r.Body.Close()

	emp.ID = id
	emps, err := a.app.UpdateEmployee(emp)
	if err != nil {
		application.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	application.RespondWithJSON(w, http.StatusOK, emps)
}

func (a *EmpHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// set header needed
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&emp); err != nil {
		application.RespondWithError(w, http.StatusBadRequest, "Invalid request")
	}
	defer r.Body.Close()

	errs := model.EmptyName(&emp)
	if errs != nil {
		application.RespondWithError(w, http.StatusBadRequest, "missing name")
		return
	}
	e := model.EmptyRole(&emp)
	if e != nil {
		application.RespondWithError(w, http.StatusBadRequest, "missing role")
		return
	}

	id, err := a.app.CreatingEmployeeDb(emp)
	if err != nil {
		application.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	emp.ID = id

	rabbit.MakeAppRb(&emp)
	application.RespondWithJSON(w, http.StatusCreated, emp)
}
