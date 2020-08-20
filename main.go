package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var assistants map[int]Assistant = make(map[int]Assistant)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/assistants", getAllAssistants)
	e.POST("/assistants", createAssistant)
	e.PUT("/assistants/:aID", updateAssistant)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

type Assistant struct {
	Name string
}

func getAllAssistants(c echo.Context) error {
	return c.JSON(http.StatusOK, assistants)
}

func createAssistant(c echo.Context) error {
	body := &Assistant{}
	if err := c.Bind(body); err != nil {
		return err
	}
	assistants[len(assistants)] = *body
	return c.JSON(http.StatusCreated, body)
}

func updateAssistant(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("aID"))
	if err != nil {
		return err
	}
	body := &Assistant{}
	if err = c.Bind(body); err != nil {
		return err
	}
	assistants[id] = *body
	return c.NoContent(http.StatusNoContent)
}
