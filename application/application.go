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
func Start(handleGet echo.HandlerFunc) {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
		AllowHeaders: []string{echo.MIMEApplicationJSON, echo.MIMEApplicationJSONCharsetUTF8},
	}))
	e.GET("/", handleGet)
	e.Logger.Fatal(e.Start(":" + getPortNumber("5000")))
}
