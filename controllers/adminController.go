package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/middleware"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAdminsController(c echo.Context) error {
	var admins []models.Admins

	if err := database.DB.Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all admins",
		"admins":  admins,
	})
}

func GetAdminByID(c echo.Context) error {
	var admins models.Admins
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&admins).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get admins ",
		"admins":  admins,
	})
}

func CreateAdminController(e echo.Context) error {
	admin := models.Admins{}
	e.Bind(&admin)

	if err := database.DB.Save(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan admins",
		"admin":   admin,
	})
}

func UpdateAdminByID(e echo.Context) error {
	admin := models.Admins{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&admin)
	admin.ID = id

	if err := database.DB.Where("id= ?", admin.ID).Updates(&admin).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data admins",
		"admin":   admin,
	})
}

func DeleteAdminByID(e echo.Context) error {
	var admin models.Admins
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&admin)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data Berhasil Dihapus",
	})
}

func LoginAdminController(e echo.Context) error {
	admin := models.Admins{}
	e.Bind(&admin)
	err := database.DB.Where("email = ? AND password = ?", admin.Email, admin.Password).First(&admin).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(admin.ID, admin.Name)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed login",
			"error":   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"token":   token,
	})
}
