package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/ogabrielrodrigues/hackaton-minerva/server/docs"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/instance"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
)

func RegisterRoutes(rt *chi.Mux) {
	rt.Mount("/api/v1/swagger", httpSwagger.WrapHandler)
	rt.Mount("/api/v1/employee", employeeRoutes())
	rt.Mount("/api/v1/feedback", feedbackRoutes())
	rt.Mount("/api/v1/answer", answersRoutes())
}

func employeeRoutes() http.Handler {
	er := chi.NewRouter()
	ei := instance.NewEmployeeInstance()

	er.Post("/", ei.Create)
	er.Get("/list", ei.List)
	er.Get("/{registry}", ei.FindByRegistry)
	er.Post("/auth", ei.Authenticate)

	return er
}

func feedbackRoutes() http.Handler {
	fr := chi.NewRouter()
	fi := instance.NewFeedbackInstance()

	fr.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/", fi.Create)
		r.Get("/", fi.List)
		r.Get("/filter", fi.Filter)
	})

	return fr
}

func answersRoutes() http.Handler {
	ar := chi.NewRouter()
	ai := instance.NewAnswerInstance()

	ar.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Get("/", ai.List)
		r.Post("/reply", ai.Reply)
	})

	return ar
}
