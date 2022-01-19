package main

import (
	"sample-crud/handlers"
	"sample-crud/stores"

	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	store := stores.New()
	handler := handlers.New(store)

	k := gofr.New()

	k.GET("/employee/{id:[0-9]+}", handler.Find)
	k.POST("/employee", handler.Create)
	k.PUT("/employee/{id:[0-9]+}", handler.Update)
	k.DELETE("/employee/{id:[0-9]+}", handler.Delete)

	k.Start()
}
