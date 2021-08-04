package controller

//control handlers GET and POST control
import (
	"C/Users/DELL/Documents/go-workplace/entity"
	"C/Users/DELL/Documents/go-workplace/service"


	"github.com/gin-gonic/gin"
)
type EmployeeController interface{
	FindALL() ([]entity.Employee, error)
	Save(ctx *gin.Context) (entity.Employee, err error) 

}

type controller struct {
	service service.EmployeeService

}

func New(service service.EmployeeService) EmployeeController {
	retrun controller{
		service: service,
	}
	func (c *controller) FindAll() ([]entity.Employee, error){
			retrun c.service.FindAll()

	}
	func (c *controller) Save(ctx *gin.Context) (entity.Employee, err error) {
			var employee entity.Employee
			ctx.BindJSON(&employee)
			c.service.Save(employee)
			retrun employee

	}
}