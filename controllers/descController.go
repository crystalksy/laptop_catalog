package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetDescs(c echo.Context) error {
	var desc []models.Descs

	if err := database.DB.Find(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Deskripsi Laptop",
		"descs":   desc,
	})
}

func GetDescByID(c echo.Context) error {
	var desc models.Descs
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Deskripsi Laptop ",
		"descs":   desc,
	})
}

func CreateDesc(e echo.Context) error {
	desc := models.Descs{}
	e.Bind(&desc)

	if err := database.DB.Save(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Deskripsi Laptop",
		"desc":    desc,
	})
}

func UpdateDescByID(e echo.Context) error {
	desc := models.Descs{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&desc)
	desc.ID = id

	if err := database.DB.Where("id= ?", desc.ID).Updates(&desc).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Deskripsi Laptop",
		"descs":   desc,
	})
}

func DeleteDescByID(e echo.Context) error {
	var desc models.Descs
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&desc)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"descs":   desc,
		"message": "Data Deskripsi Laptop Berhasil Dihapus",
	})
}
