package entity

type EmployeeDTO struct {
	Id   int    `json:"id,string,omitempty"`
	Name string `json:"name"`
}
