package stores

import (
	"sample-crud/db"
	"sample-crud/models"

	"developer.zopsmart.com/go/gofr/pkg/errors"
)

type employee struct {
	db []models.Employee
}

func New() *employee {
	return &employee{db: db.Employees}
}

func (s *employee) Find(id string) (models.Employee, error) {
	for _, Employee := range s.db {
		if Employee.ID == id {
			return Employee, nil
		}
	}

	return models.Employee{}, errors.EntityNotFound{Entity: "Employee", ID: id}
}

func (s *employee) Create(employee models.Employee) error {
	_, err := s.Find(employee.ID)
	if err == nil {
		return errors.EntityAlreadyExists{}
	}
	s.db = append(s.db, employee)

	return nil
}

func (s *employee) Update(id string, employee models.Employee) error {
	for i := range s.db {
		if s.db[i].ID == id {
			s.db[i].CTC = employee.CTC
			s.db[i].Name = employee.Name
			s.db[i].Email = employee.Email
			return nil
		}
	}

	return errors.EntityNotFound{Entity: "Employee", ID: id}
}

func (s *employee) Delete(id string) error {
	for i, Employee := range s.db {
		if Employee.ID == id {
			s.db = append(s.db[:i], s.db[i+1:]...)
			return nil
		}
	}

	return errors.EntityNotFound{Entity: "Employee", ID: id}
}
