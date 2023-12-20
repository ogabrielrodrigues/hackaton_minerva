package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/database"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/logger"
)

type pg struct{}

func GetDB() database.DB {
	return &pg{}
}

func (*pg) Connect() (*sql.DB, error) {
	conn, err := sql.Open("postgres", config.GetConfig().Database)
	if err != nil {
		logger.Err("error initializing database", err)
	}

	err = conn.Ping()

	return conn, err
}
