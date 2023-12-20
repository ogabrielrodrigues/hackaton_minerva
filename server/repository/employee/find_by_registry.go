package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (er *employeeRepository) FindByRegistry(registry string) (entity.Employee, *rest.RestErr) {
	conn, err := er.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM "tb_employees" WHERE registry=$1`, registry)
	if row.Err() != nil {
		return nil, rest.NewNotFoundErr(err.Error())
	}

	var name, email, sector, unit string
	var administrator bool

	if err := row.Scan(&registry,
		&name,
		&email,
		&sector,
		&unit,
		&administrator); err != nil {
		return nil, rest.NewNotFoundErr(err.Error())
	}

	employee := entity.NewEmployee(name, email, sector, unit, administrator)
	employee.SetRegistry(registry)

	return employee, nil
}
