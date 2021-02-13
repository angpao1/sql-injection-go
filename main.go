package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "employee_db"
)

func main() {
	e := echo.New()

	// Routes
	e.GET("/employees/unsafe/:id", getEmployee1)
	e.GET("/employees/safe/:id", getEmployee2)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

}

func getEmployee1(c echo.Context) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	employeeID := c.Param("id")
	query := "SELECT id, name, email FROM employee WHERE id = " + employeeID
	row, err := db.QueryContext(ctx, query)
	if err != nil {
		fmt.Println(err)
	}

	var employees []Employee

	for row.Next() {
		var data Employee
		err := row.Scan(
			&data.ID,
			&data.Name,
			&data.Email)

		if err != nil {
			fmt.Println(err)
		}
		employees = append(employees, data)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, employees)
}

func getEmployee2(c echo.Context) error {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	employeeID := c.Param("id")
	query := "SELECT id, name, email FROM employee WHERE id = $1"
	row, err := db.QueryContext(ctx, query, employeeID)
	if err != nil {
		fmt.Println(err)
	}

	var employees []Employee

	for row.Next() {
		var data Employee
		err := row.Scan(
			&data.ID,
			&data.Name,
			&data.Email)

		if err != nil {
			fmt.Println(err)
		}
		employees = append(employees, data)
	}

	defer db.Close()
	return c.JSON(http.StatusOK, employees)
}

// Employee use for response
type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
