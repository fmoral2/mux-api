package model

import "errors"

type Employee struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Wage     float64 `json:"wage,omitempty"`
	Role     string  `json:"role,omitempty"`
	Location string  `json:"location,omitempty"`
}

func EmptyName(emp *Employee) error {
	if emp.Name == "" {
		return errors.New("missing name")
	}
	return nil
}

func EmptyRole(emp *Employee) error {
	if emp.Role == "" {
		return errors.New("missing role")
	}
	return nil
}
