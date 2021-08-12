package entity

func ToEmployee(employeeDTO EmployeeDTO) Employee {
	employeeDTOO := EmployeeDTO{Id: employeeDTO.Id, Name: employeeDTO.Name}
	employee := Employee(employeeDTOO)
	return employee
}

func ToEmployeeDTO(employee Employee) EmployeeDTO {
	employeee := Employee{Id: employee.Id, Name: employee.Name}
	employeeDTO := EmployeeDTO(employeee)

	return employeeDTO
}

func ToEmployeeDTOs(employees []Employee) []EmployeeDTO {
	employeedtos := make([]EmployeeDTO, len(employees))

	for i, itm := range employees {
		employeedtos[i] = ToEmployeeDTO(itm)
	}

	return employeedtos
}
