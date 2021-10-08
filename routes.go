package main

import (
	"bpk-dx/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func makeRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.Home{}.HomeMain)
	r.Get("/employee", controllers.Emp{}.EmpMain)
	r.Get("/employee/add", controllers.Emp{}.EmpAdd)
	r.Post("/employee/add", controllers.Emp{}.EmpCreate)

	r.HandleFunc("/employee/delete", controllers.Emp{}.DeleteEmployee)
	return r
}
