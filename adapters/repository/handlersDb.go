package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/morlfm/rest-api/application/model"
)

const (
	getEmp       = "SELECT id,name,location,role,wage FROM employeesDb WHERE id=$1"
	getEmployess = "SELECT * FROM employeesDb"
	updateEmp    = "UPDATE employeesDb SET name=$1,location=$2 WHERE id=$3 RETURNING id , name, location"
	deleteEmp    = "DELETE FROM employeesDb WHERE id=$1"
	createEmp    = "INSERT INTO employeesDb(id,name,location,wage,role) VALUES($1,$2,$3,$4,$5) RETURNING id"
)

type EmployeePG struct {
	db *sql.DB
}

type EmployeeDb struct {
	ID       string  `db:"id"`
	Name     string  `db:"name"`
	Wage     float64 `db:"wage"`
	Location string  `db:"location"`
	Role     string  `db:"role"`
}

func MakeRepository(db *sql.DB) *EmployeePG {
	return &EmployeePG{db: db}
}

func (r *EmployeePG) CreateEmployeeDb(name string, location string, wage float64, role string) (string, error) {
	var id string
	empID := uuid.NewString()
	err := r.db.QueryRow(createEmp, empID, name, location, wage, role).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}

func (r *EmployeePG) GetEmployeeDb(id string) (model.Employee, error) {

	e := EmployeeDb{}
	err := r.db.QueryRow(getEmp, id).Scan(&e.Name, &e.ID, &e.Location, &e.Role, &e.Location)
	emp := model.Employee{ID: e.ID, Name: e.Name, Location: e.Location, Role: e.Role, Wage: e.Wage}

	return emp, err
}

func (r *EmployeePG) GetEmployeesDb() ([]model.Employee, error) {

	rows, err := r.db.Query(getEmployess)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	emps := []model.Employee{}

	//
	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.ID, &emp.Name, &emp.Location, &emp.Wage, &emp.Role); err != nil {
			return nil, err
		}

		emps = append(emps, emp)

	}

	return emps, err
}

func (r *EmployeePG) DeleteEmployeeDb(id string) (model.Employee, error) {

	e := EmployeeDb{}

	result, err := r.db.Exec(deleteEmp, id)

	if rowsAffectd, _ := result.RowsAffected(); rowsAffectd == 0 {
		return model.Employee{}, errors.New("")
	}

	emp := model.Employee{ID: e.ID}

	return emp, err
}

func (r *EmployeePG) UpdateEmployeeDb(emp model.Employee) (model.Employee, error) {

	err := r.db.QueryRow(updateEmp, emp.Name, emp.Location, emp.ID).Scan(&emp.ID, &emp.Name, &emp.Location)

	// mapping only 3 to return 3
	res := model.Employee{ID: emp.ID, Name: emp.Name, Location: emp.Location}

	return res, err
}
