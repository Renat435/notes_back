package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/mail"
)

func (s *Service) CreateUser(c echo.Context) error {
	var user NewUser

	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("invalid email address: %s", err))
	}

	hp, err := s.HashPassword(user.Password)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	newUser, err := s.usersData.CreateUser(user.Email, hp)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, newUser)
}

func (s *Service) GetUsers(c echo.Context) error {
	c.Logger().Infof("Get users")

	return nil
}
