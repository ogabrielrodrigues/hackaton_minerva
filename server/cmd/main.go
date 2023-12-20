package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/router"
)

// @title Minerva Suggestion API
// @version 1.0
// @description Minerva Suggestion API
// @BasePath /api/v1
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}

	rt := chi.NewRouter()
	router.RegisterMiddlewares(rt)
	router.RegisterRoutes(rt)

	logger.Info(fmt.Sprintf("server running on %s", config.GetConfig().Port))
	if err := http.ListenAndServe(config.GetConfig().Port, rt); err != nil {
		logger.Err("error initializing server", err)
	}
}
