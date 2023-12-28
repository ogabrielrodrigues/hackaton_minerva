package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (er *employeeRepository) Authenticate(employee entity.Employee) (entity.Employee, *rest.RestErr) {
	conn, err := er.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`SELECT * FROM "tb_employees" 
		WHERE email=$1 AND password=$2`,
		employee.GetEmail(),
		employee.GetPassword(),
	)

	var registry, name, email, sector, unit, password string
	var administrator bool

	if err := row.Scan(
		&registry,
		&name,
		&email,
		&sector,
		&unit,
		&administrator,
		&password,
	); err != nil {
		return nil, rest.NewForbiddenErr("email or password are incorrect")
	}

	employee.SetRegistry(registry)
	employee.SetName(name)
	employee.SetSector(sector)
	employee.SetUnit(unit)
	employee.SetAdministrator(administrator)

	return employee, nil
}
