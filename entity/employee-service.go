package entity

type EmployeeService struct {
	EmployeeRepository EmployeeRepository
}

func GetEmployeeService(e EmployeeRepository) EmployeeService {
	return EmployeeService{EmployeeRepository: e}
}

func (stmt *EmployeeService) FindAll() ([]Employee, error) {
	// var employees []Employee
	return stmt.EmployeeRepository.FindAll()
	// return employees, nil
}

func (stmt *EmployeeService) FindByID(id int) (Employee, error) {
	// var employees Employee
	return stmt.EmployeeRepository.FindByID(id)
	// return employees, nil
}

func (stmt *EmployeeService) Save(employee Employee) (Employee, error) {
	//var employees Employee
	return stmt.EmployeeRepository.Save(employee)
	//return employees, nil
}

func (p *EmployeeService) Delete(employee Employee) error {
	p.EmployeeRepository.Delete(employee)

	return nil
}
