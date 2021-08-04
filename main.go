package main

import (
	//"database/sql"
	//"fmt"
	//"log"
	"C/Users/DELL/Documents/go-workplace/entity"
	"C\Users\DELL\Documents\go-workplace\service"
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"
	//"gorm.io/gorm"
)

var (
	employeeService service.EmployeeService= service.New()
	EmployeeController controller.Controller = controller.New(employeeService)

)

func main() {
	server := gin.Default()

	server.GET("/employees", func(ctx *gin.Context){
		ctx.JSON(200,EmployeeController.FindALL())
		})
	server.POST("/employees", func(ctx *gin.Context){
		ctx.JSON(200,EmployeeController.Save(ctx))
		})
	server.Run(":8080")

}
