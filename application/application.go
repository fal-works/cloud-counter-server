package application

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func getPortNumber(defaultNumber string) string {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		return defaultNumber
	}
	return portNumber
}

// Start starts the web application.
func Start(handleGet echo.HandlerFunc, handleIncrement echo.HandlerFunc) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"https://preview.openprocessing.org", "https://fal-works.github.io/", "https://www.fal-works.com"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
		AllowHeaders: []string{echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8},
	}))
	e.GET("/", handleGet)
	e.GET("/increment", handleIncrement)
	e.Logger.Fatal(e.Start(":" + getPortNumber("5000")))
}
