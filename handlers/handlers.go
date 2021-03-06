package handlers

import (
	"sample-crud/models"
	"sample-crud/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

type employee struct {
	store stores.Employee
}

func New(store stores.Employee) *employee {
	return &employee{store: store}
}

func (s *employee) Find(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")
	return s.store.Find(id)
}

func (s *employee) Create(c *gofr.Context) (interface{}, error) {
	var employee models.Employee

	if err := c.Bind(employee); err != nil {
		return nil, err
	}

	return s.store.Create(employee), nil
}

func (s *employee) Update(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	var employee models.Employee

	if err := c.Bind(employee); err != nil {
		return nil, err
	}

	return s.store.Update(id, employee), nil
}

func (s *employee) Delete(c *gofr.Context) (interface{}, error) {
	id := c.PathParam("id")

	return s.store.Delete(id), nil
}
