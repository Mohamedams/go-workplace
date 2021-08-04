package controller

//control handlers GET and POST control
import (
	"bin/Documents/go-workplace/entity"
	"bin/Documents/go-workplace/service"

	"github.com/gin-gonic/gin"
)

type EmployeeController interface {
	Save(ctx *gin.Context) entity.Employee
	FindALL() []entity.Employee
}

type controller struct {
	service service.EmployeeService
}

func NewController(service service.EmployeeService) EmployeeController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindALL() []entity.Employee {
	return c.service.FindALL()

}
func (c *controller) Save(ctx *gin.Context) entity.Employee {
	var employee entity.Employee
	ctx.BindJSON(&employee)
	c.service.Save(employee)
	return employee

}
