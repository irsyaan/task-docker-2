package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Use the logger middleware to log all requests
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route to display plain text
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Someone hit me!")
		return c.String(http.StatusOK, "The only thing we never get enough of is love; and the only thing we never give enough of is love.")
	})

	// Start the server on port 80
	port := "80"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	e.Logger.Fatal(e.Start(":" + port))
}
