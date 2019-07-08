package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func getPortNumber(defaultNumber string) string {
	portNumber := os.Getenv("PORT")
	if portNumber == "" {
		return defaultNumber
	}
	return portNumber
}

type countValue struct {
	Count int `json:"count"`
}

func getCountJSON(c echo.Context) error {
	responseData := &countValue{
		Count: 1,
	}
	return c.JSON(http.StatusOK, responseData)
}

func main() {
	e := echo.New()
	e.GET("/", getCountJSON)
	e.Logger.Fatal(e.Start(":" + getPortNumber("5000")))
}
