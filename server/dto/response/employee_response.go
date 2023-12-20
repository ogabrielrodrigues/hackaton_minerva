package response

type EmployeeResponse struct {
	Registry      string `json:"registry"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Sector        string `json:"sector"`
	Unit          string `json:"unit"`
	Administrator bool   `json:"administrator"`
}
