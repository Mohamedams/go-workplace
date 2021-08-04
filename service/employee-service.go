package service

import (
	"bin/Documents/go-workplace/entity"
)

type EmployeeService interface {
	Save(entity.Employee) entity.Employee
	FindALL() []entity.Employee
}

type employeeService struct {
	employees []entity.Employee
}

func New() EmployeeService {
	return &employeeService{}
}

func (service *employeeService) Save(employee entity.Employee) entity.Employee {
	service.employees = append(service.employees, employee)

	return employee

}
func (service *employeeService) FindALL() []entity.Employee {

	return service.employees

}
