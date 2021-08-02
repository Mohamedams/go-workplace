package main

import (
	//"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Employee STRUCT
type Employee struct {
	ID   int
	Name string
}

//------------ CONNECTION NAME------------
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Abouras"
	dbname   = "postgres"
)

func main() {

	//----------------------------------------------------------------OPEN CONNECTION ------------------------------
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})

	if err != nil {

		log.Fatal(err)
	}

	//---------------------------------------------------------------------------------------------------------------------------------------
	//------------CRUD Starts Here----------------------------------------------------------------
	//..

	Employee1 := Employee{ID: 4, Name: "Mohamed"}
	var Employee2 []Employee
	var Employee3 []Employee
	var Employee6 []Employee
	result := db.Create(&Employee1)
	result4 := db.Find(&Employee2)
	result2 := db.Last(&Employee3)
	result5 := db.Order("name").Find(&Employee6)

	if result.Error != nil {
		log.Fatal(result.Error)
	}
	if result5.Error != nil {
		log.Fatal(result5.Error)
	}
	if result4.Error != nil {
		log.Fatal(result4.Error)
	}
	if result2.Error != nil {
		log.Fatal(result.Error)
	}

	var Employee4 Employee
	db.First(&Employee4)

	Employee4.Name = "Mohamed"
	Employee4.ID = 3
	db.Save(&Employee4)

	db.Model(&Employee4).Update("name", "hello")

	db.Delete(&Employee{}, 1)

	//------------
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(500, gin.H{
			"message": &Employee6,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	//------------
	// r.GET("/create", func(c *gin.Context) {
	// 	Id := c.Param("id")

	// 	Name := c.Param("name")

	// Employee1 := Employee{ID: Id, Name: Name}
	// 	result1 := db.Create(&Employee1)
	// 	if result1.Error != nil {
	// 		log.Fatal(result1.Error)
	// 	}

	// 	})

	// 	fmt.Sprintln(Id, Name)
	// })

}
