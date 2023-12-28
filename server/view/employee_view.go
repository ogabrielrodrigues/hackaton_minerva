package view

import (
	"github.com/ogabrielrodrigues/hackaton-minerva/server/dto/response"
	"github.com/ogabrielrodrigues/hackaton-minerva/server/entity"
)

func EmployeeToView(employee entity.Employee) *response.EmployeeResponse {
	var feedbacks []response.FeedbackResponse

	for _, feedback := range employee.GetFeedbacks() {
		feedbacks = append(feedbacks, *FeedbackToView(feedback))
	}

	return &response.EmployeeResponse{
		Registry:      employee.GetRegistry(),
		Name:          employee.GetName(),
		Email:         employee.GetEmail(),
		Sector:        employee.GetSector(),
		Unit:          employee.GetUnit(),
		Administrator: employee.GetAdministrator(),
		Feedbacks:     feedbacks,
	}
}
