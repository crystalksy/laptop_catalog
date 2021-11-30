package controllers

//import packages
import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data feature
func GetFeatures(c echo.Context) error {
	var feature []models.Features

	if err := database.DB.Find(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Features Laptop",
		"feature": feature,
	})
}

//Fungsi get feature by ID
func GetFeatureByID(c echo.Context) error {
	var feature models.Features
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Features Laptop ",
		"feature": feature,
	})
}

//fungsi create new feature
func CreateFeature(e echo.Context) error {
	feature := models.Features{}
	e.Bind(&feature)

	if err := database.DB.Save(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Feature Laptop",
		"feature": feature,
	})
}

//Fungsi Update Tabel Feature
func UpdateFeatureByID(e echo.Context) error {
	feature := models.Features{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&feature)
	feature.ID = id

	if err := database.DB.Where("id= ?", feature.ID).Updates(&feature).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Feature Laptop",
		"feature": feature,
	})
}

//Fungsi hapus Features data
func DeleteFeatureByID(e echo.Context) error {
	var feature models.Features
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&feature)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"feature": feature,
		"message": "Data Feature Berhasil Dihapus",
	})
}
