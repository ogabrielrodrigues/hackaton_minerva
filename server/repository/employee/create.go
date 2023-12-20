package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (er *employeeRepository) Create(employee entity.Employee) (entity.Employee, *rest.RestErr) {
	conn, err := er.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`INSERT INTO "tb_employees" (name, email, sector, unit, administrator) VALUES ($1, $2, $3, $4, $5) RETURNING registry`,
		employee.GetName(),
		employee.GetEmail(),
		employee.GetSector(),
		employee.GetUnit(),
		employee.GetAdministrator(),
	)
	if row.Err() != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var registry string
	if err := row.Scan(&registry); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	employee.SetRegistry(registry)
	return employee, nil
}
