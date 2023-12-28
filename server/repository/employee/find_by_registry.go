package repository

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/view"
)

func (er *employeeRepository) FindByRegistry(registry string) (entity.Employee, *rest.RestErr) {
	conn, err := er.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`SELECT
			e.*, 
			JSON_AGG(f.*) AS feedbacks
		FROM "tb_employees" e
		LEFT JOIN "tb_feedbacks" f
		ON e.registry = f.employee_registry
		WHERE e.registry = $1
		GROUP BY e.registry`, registry)

	var name, email, sector, unit, password string
	var administrator bool
	var json_arr []byte

	if err := row.Scan(&registry,
		&name,
		&email,
		&sector,
		&unit,
		&administrator,
		&password,
		&json_arr); err != nil {
		return nil, rest.NewNotFoundErr(err.Error())
	}

	employee := entity.NewEmployee(name, email, sector, unit, password, administrator)
	employee.SetRegistry(registry)
	employee.SetPassword(password)

	if string(json_arr) != "[null]" {
		feedbacks := view.ViewToFeedback(json_arr)

		employee.SetFeedbacks(feedbacks)
	}

	return employee, nil
}
