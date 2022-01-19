package stores

import "sample-crud/models"

type Employee interface {
	Find(id string) (models.Employee, error)
	Create(Employee models.Employee) error
	Update(id string, Employee models.Employee) error
	Delete(id string) error
}