package main

// import (
// 	"github.com/gin-gonic/gin"
// 	//"gorm.io/driver/postgres"
// 	//"gorm.io/gorm"
// )

// var (
// 	employeeService    service.EmployeeService       = service.New()
// 	EmployeeController controller.EmployeeController = controller.NewController(employeeService)
// )

// func main() {
// 	server := gin.Default()

// 	server.GET("/employees", func(ctx *gin.Context) {
// 		ctx.JSON(200, EmployeeController.FindALL())
// 	})
// 	server.POST("/employees", func(ctx *gin.Context) {
// 		ctx.JSON(200, EmployeeController.Save(ctx))
// 	})
// 	server.Run(":8080")

// }

import (
	"fmt"
	"log"

	//"os"

	"bin/Documents/go-workplace/entity"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Abouras"
	dbname   = "postgres"
)

func initDB() *gorm.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Employee{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	employeeDB, err := entity.ProvideEmployeeRepostiory(db)
	if err != nil {
		log.Fatal(err)
	}
	employeeService := entity.GetEmployeeService(employeeDB)

	employeeAPI := entity.GetEmployeeAPI(employeeService)

	r := gin.Default()

	r.GET("/employees", employeeAPI.FindAll)
	r.GET("/employees/:id", employeeAPI.FindByID)
	r.POST("/employees", employeeAPI.Create)
	r.PUT("/employees/:id", employeeAPI.Update)
	r.DELETE("/employees/delete/:id", employeeAPI.Delete)

	error := r.Run()
	if error != nil {
		panic(err)
	}
}
