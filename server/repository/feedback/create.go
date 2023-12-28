package repository

import (
	"time"

	"github.com/ogabrielrodrigues/hackaton-minerva/server/config/rest"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func (fr *feedbackRepository) Create(feedback entity.Feedback) (entity.Feedback, *rest.RestErr) {
	conn, err := fr.database.Connect()
	if err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}
	defer conn.Close()

	row := conn.QueryRow(
		`INSERT INTO "tb_feedbacks" (employee_registry, content) VALUES ($1, $2) RETURNING id, sent_at`,
		feedback.GetEmployeeRegistry(),
		feedback.GetContent(),
	)

	var id string
	var sent_at time.Time
	if err := row.Scan(&id, &sent_at); err != nil {
		return nil, rest.NewInternalServerErr(err.Error())
	}

	feedback.SetID(id)
	feedback.SetSentAt(sent_at)

	return feedback, nil
}
