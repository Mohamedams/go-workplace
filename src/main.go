package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Employee struct {
	Id   int
	name string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Abouras"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)

	if err != nil {

		log.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt := "CREATE TABLE IF NOT EXISTS Employees(Id SERIAL PRIMARY KEY , Name VARCHAR(255) NOT NULL);"

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
	sqlStatement := `
INSERT INTO Employees(name) 
VALUES ('mohamed')`
	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
	sqlStatementDelete := ` 
	DELETE FROM Employees
WHERE Id=1`
	_, err = db.Exec(sqlStatementDelete)
	if err != nil {
		panic(err)
	}

	sqlStatementUpdate := ` 
	UPDATE Employees
	SET name ='MOHAMED'
WHERE Id=2`
	_, err = db.Exec(sqlStatementUpdate)
	if err != nil {
		panic(err)
	}
	var myEmployee Employee
	userSql := "SELECT * FROM Employees WHERE id = 5"

	err = db.QueryRow(userSql).Scan(&myEmployee.Id, &myEmployee.name)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	fmt.Printf("Hi %s, welcome back!\n", myEmployee.name)
}
