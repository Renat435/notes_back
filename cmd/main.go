package main

import (
	"github.com/labstack/echo/v4/middleware"
	"notes/internal/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := PostgresConnection()

	e := echo.New()

	e.Use(middleware.Logger())

	if err != nil {
		e.Logger.Fatal(err)
	}

	api := e.Group("/api")

	svc := service.InitServices(db)

	api.POST("/register", svc.CreateUser)
	api.POST("/auth", svc.GetUsers)

	e.Logger.Fatal(e.Start(":1323"))
}
