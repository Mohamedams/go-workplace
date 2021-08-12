package entity

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func ProvideEmployeeRepostiory(DB *gorm.DB) (EmployeeRepository, error) {
	return EmployeeRepository{DB: DB}, nil
}

func (stmt *EmployeeRepository) FindAll() ([]Employee, error) {
	var employees []Employee
	stmt.DB.Find(&employees)
	return employees, nil
}

func (stmt *EmployeeRepository) FindByID(id int) (Employee, error) {
	var employees Employee
	stmt.DB.First(&employees, id)
	return employees, nil
}

func (stmt *EmployeeRepository) Save(employee Employee) (Employee, error) {
	var employees Employee
	stmt.DB.Save(&employee)
	return employees, nil
}

func (e *EmployeeRepository) Delete(employee Employee) error {
	e.DB.Delete(&employee)

	return nil
}
