package request

type CreateEmployeeRequest struct {
	Name          string `json:"name" validate:"required,min=4,max=100"`
	Email         string `json:"email" validate:"required,email"`
	Sector        string `json:"sector" validate:"required,min=4,max=100"`
	Unit          string `json:"unit" validate:"required,min=4,max=100"`
	Administrator bool   `json:"administrator"`
}
