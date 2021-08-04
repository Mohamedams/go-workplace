package service

import (
	"C/Users/DELL/Documents/go-workplace/entity"
	
)
type EmployeeService interface{
	Save(entity.Employee) (entity.Employee,error)
	FindALL() []entity.Employee
}

type employeeService struct {
	employees []entity.Employee
}
 
func New() EmployeeService {
	retrun &emplyeesService{}
}

func (service *employeeService) Save(employee entity.Employee) entity.Employee{
	service.employees,err = append(service.employees, employee)
	if err != nil {
		log.Printf("Error saving employee")
	}
	retrun (employee,"no errors")

}
func (service *employeeService) FindALL() []entity.Employee{

	retrun service.employees


}