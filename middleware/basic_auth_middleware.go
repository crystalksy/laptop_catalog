package middleware

import (
	"laptop_catalog/database"
	"laptop_catalog/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(adminname, password string, c echo.Context) (bool, error) {

	var admin models.Admins

	err := database.DB.Where("email = ? AND password = ?", adminname, password).First(&admin).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
