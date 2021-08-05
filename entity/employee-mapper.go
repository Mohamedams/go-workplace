package entity

func ToEmployee(employeeDTO EmployeeDTO) Employee {
	return Employee{Name: employeeDTO.Name}
}

func ToEmployeeDTO(employee Employee) EmployeeDTO {
	return EmployeeDTO{Id: employee.Id, Name: employee.Name}
}

func ToEmployeeDTOs(employees []Employee) []EmployeeDTO {
	employeedtos := make([]EmployeeDTO, len(employees))

	for i, itm := range employees {
		employeedtos[i] = ToEmployeeDTO(itm)
	}

	return employeedtos
}
