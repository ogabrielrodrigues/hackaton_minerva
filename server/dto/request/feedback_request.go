package request

type CreateFeedbackRequest struct {
	Employee_registry string `json:"employee_registry" validate:"required,uuid4"`
	Content           string `json:"content" validate:"required"`
}
