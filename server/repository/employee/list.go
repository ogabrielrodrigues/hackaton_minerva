package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (er *employeeRepository) List() ([]entity.Employee, *rest.RestErr) {
	conn, err := er.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM "tb_employees"`)
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var registry, name, email, sector, unit, password string
	var administrator bool

	var employees []entity.Employee

	for rows.Next() {
		defer rows.Close()

		if err := rows.Scan(
			&registry,
			&name,
			&email,
			&sector,
			&unit,
			&administrator,
			&password,
		); err != nil {
			return nil, rest.NewInternalServerErr(err.Error())
		}

		employee := entity.NewEmployee(name, email, sector, unit, password, administrator)
		employee.SetRegistry(registry)
		employee.SetPassword(password)

		employees = append(employees, employee)
	}

	return employees, nil
}
