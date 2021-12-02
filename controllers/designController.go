package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDesigns(c echo.Context) error {
	var design []models.Designs

	if err := database.DB.Find(&design).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Design Laptop",
		"design":  design,
	})
}

func GetDesignByID(c echo.Context) error {
	var design models.Designs
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&design).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Data Design Laptop ",
		"design":  design,
	})
}

func CreateDesign(e echo.Context) error {
	design := models.Designs{}
	e.Bind(&design)

	if err := database.DB.Save(&design).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Design Laptop",
		"design":  design,
	})
}

func UpdateDesignByID(e echo.Context) error {
	design := models.Designs{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&design)
	design.ID = id

	if err := database.DB.Where("id= ?", design.ID).Updates(&design).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Design Laptop",
		"design":  design,
	})
}

func DeleteDesignByID(e echo.Context) error {
	var design models.Designs
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&design)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"design":  design,
		"message": "Data Design Berhasil Dihapus",
	})
}
