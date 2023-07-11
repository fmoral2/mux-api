package repository

import (
	"database/sql"
	"errors"
	"log"
	"math"

	"github.com/fmoral2/mux-api/application/model"
	"github.com/google/uuid"
)

const (
	getEmp       = "SELECT id,name,location,role,wage FROM employeesDb WHERE id=$1"
	getEmployess = "SELECT id,name,location,wage,role  FROM employeesDb ORDER BY id ASC  OFFSET $1 LIMIT $2 "
	getAll       = "SELECT id,name,location,wage,role FROM employeesDb"
	updateEmp    = "UPDATE employeesDb SET name=$1,location=$2 WHERE id=$3 RETURNING id , name, location"
	deleteEmp    = "DELETE FROM employeesDb WHERE id=$1"
	createEmp    = "INSERT INTO employeesDb(id,name,location,wage,role) VALUES($1,$2,$3,$4,$5) RETURNING id"
	getEmpsCount = "SELECT COUNT(*) FROM employeesDb"
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
	err := r.db.QueryRow(getEmp, id).Scan(&e.ID, &e.Name, &e.Location, &e.Role, &e.Wage)
	emp := model.Employee{ID: e.ID, Name: e.Name, Location: e.Location, Role: e.Role, Wage: e.Wage}

	return emp, err
}

func (r *EmployeePG) GetEmployeesDb() ([]model.Employee, error) {

	rows, err := r.db.Query(getAll)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var emps []model.Employee

	for rows.Next() {
		var emp model.Employee
		if err = rows.Scan(&emp.ID, &emp.Name, &emp.Role, &emp.Wage, &emp.Location); err != nil {
			return nil, err
		}

		emps = append(emps, emp)

	}
	return emps, nil

}

func (r *EmployeePG) GetEmployeesFilterDb(page int64, size int64) (interface{}, error) {
	if size == 0 {
		return model.PageResponse{}, nil
	}

	limit := size
	offSet := page * size

	rows, err := r.db.Query(getEmployess, offSet, limit)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var emps []model.Employee

	for rows.Next() {
		var emp model.Employee
		if err = rows.Scan(&emp.ID, &emp.Name, &emp.Role, &emp.Wage, &emp.Location); err != nil {
			return nil, err
		}

		emps = append(emps, emp)

	}
	var count int64
	countRows, err := r.db.Query(getEmpsCount)
	if err != nil {
		return nil, err
	}
	for countRows.Next() {
		if err = countRows.Scan(&count); err != nil {
			return count, nil
		}

	}
	employeesResponse := model.EmployeesResponse{
		Items: emps,
		PageResponse: model.PageResponse{
			Page:       page,
			TotalItems: count,
			TotalPages: int64(math.Ceil(float64(count) / float64(size))),
		},
	}

	return employeesResponse, err
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
