package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
)

func main() {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"user",
		"password",
		"db",
		"3306",
		"sampledb",
	)
	conn, err := gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error)
	}

	err = conn.DB().Ping()
	if err != nil {
		panic(err)
	}

	conn.LogMode(true)
	conn.Set("gorm:table_options", "ENGINE=InnoDB")

	e := echo.New()

	// Routes
	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":8080"))
}


// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")	
}
