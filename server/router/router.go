package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/ogabrielrodrigues/hackaton-minerva/server/docs"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/instance"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(rt *chi.Mux) {
	rt.Mount("/api/v1/swagger", httpSwagger.WrapHandler)
	rt.Mount("/api/v1/employee", employeeRoutes())
}

func employeeRoutes() http.Handler {
	er := chi.NewRouter()
	ei := instance.NewEmployeeInstance()

	er.Post("/", ei.Create)
	er.Get("/list", ei.List)
	er.Get("/{registry}", ei.FindByRegistry)

	return er
}
