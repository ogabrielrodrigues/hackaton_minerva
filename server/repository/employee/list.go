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

	row, err := conn.Query(`SELECT * FROM "tb_employees"`)
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	var registry, name, email, sector, unit string
	var administrator bool

	var employees []entity.Employee

	for row.Next() {
		if row.Err() != nil {
			return nil, rest.NewInternalServerErr(err.Error())
		}

		if err := row.Scan(
			&registry,
			&name,
			&email,
			&sector,
			&unit,
			&administrator,
		); err != nil {
			return nil, rest.NewInternalServerErr(err.Error())
		}

		employee := entity.NewEmployee(name, email, sector, unit, administrator)
		employee.SetRegistry(registry)

		employees = append(employees, employee)
	}

	return employees, nil
}
