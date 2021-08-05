package entity

import (
	"strconv"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmployeeAPI struct {
	EmployeeService EmployeeService
}

func GetEmployeeAPI(e EmployeeService) EmployeeAPI {
	return EmployeeAPI{EmployeeService: e}
}

func (e *EmployeeAPI) FindAll(c *gin.Context) {
	employees, err := e.EmployeeService.FindAll()

	if err != nil {
		log.Printf("No Employees or Error finding some")
	}

	c.JSON(http.StatusOK, gin.H{"employees": ToEmployeeDTOs(employees)})

}
func (e *EmployeeAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := e.EmployeeService.FindByID(int(id))
	if err != nil {
		log.Printf("EmployeeService.FindByID")
	}

	c.JSON(http.StatusOK, gin.H{"product": ToEmployeeDTO(product)})
}
func (e *EmployeeAPI) Create(c *gin.Context) {
	var employeeDTO EmployeeDTO
	err := c.BindJSON(&employeeDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdEmployee, err := e.EmployeeService.Save(ToEmployee(employeeDTO))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"employee": ToEmployeeDTO(createdEmployee)})
}

func (e *EmployeeAPI) Update(c *gin.Context) {
	var employeeDTO EmployeeDTO
	err := c.BindJSON(&employeeDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := e.EmployeeService.FindByID(int(id))

	employee.Name = employeeDTO.Name
	if err != nil {
		log.Fatalln(err)
	}
	if employee == (Employee{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	e.EmployeeService.Save(employee)

	c.Status(http.StatusOK)
}

func (e *EmployeeAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := e.EmployeeService.FindByID(int(id))
	if err != nil {
		log.Fatal(err)
	}
	if product == (Employee{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	e.EmployeeService.Delete(product)

	c.Status(http.StatusOK)
}
