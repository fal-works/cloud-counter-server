package main

import (
	"log"
	"net/http"

	"github.com/fal-works/cloud-counter/application"
	"github.com/fal-works/cloud-counter/database"
	_ "github.com/lib/pq"

	"github.com/labstack/echo"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	getCountJSON := func(echoContext echo.Context) error {
		countInstance := database.GetCount(db)
		return echoContext.JSON(http.StatusOK, countInstance)
	}

	getIncrementedCountJSON := func(echoContext echo.Context) error {
		countInstance := database.GetIncrementedCount(db)
		return echoContext.JSON(http.StatusOK, countInstance)
	}

	application.Start(getCountJSON, getIncrementedCountJSON)
}
