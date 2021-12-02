package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCameras(c echo.Context) error {
	var camera []models.Cameras

	if err := database.DB.Find(&camera).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Kamera Laptop",
		"camera":  camera,
	})
}

func GetCameraByID(c echo.Context) error {
	var camera models.Cameras
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&camera).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Data Kamera Laptop ",
		"camera":  camera,
	})
}

func CreateCamera(e echo.Context) error {
	camera := models.Cameras{}
	e.Bind(&camera)

	if err := database.DB.Save(&camera).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Kamera",
		"camera":  camera,
	})
}

func UpdateCameraByID(e echo.Context) error {
	camera := models.Cameras{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&camera)
	camera.ID = id

	if err := database.DB.Where("id= ?", camera.ID).Updates(&camera).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Kamera",
		"camera":  camera,
	})
}

func DeleteCameraByID(e echo.Context) error {
	var camera models.Cameras
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&camera)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"camera":  camera,
		"message": "Data Camera Berhasil Dihapus",
	})
}
