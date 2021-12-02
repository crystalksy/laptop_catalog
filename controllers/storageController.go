package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetStorages(c echo.Context) error {
	var storage []models.Storages

	if err := database.DB.Find(&storage).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Storage Laptop",
		"storage": storage,
	})
}

func GetStorageByID(c echo.Context) error {
	var storage models.Storages
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&storage).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Storage Laptop ",
		"storage": storage,
	})
}

func CreateStorage(e echo.Context) error {
	storage := models.Storages{}
	e.Bind(&storage)

	if err := database.DB.Save(&storage).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Storage Laptop",
		"storage": storage,
	})
}

func UpdateStorageByID(e echo.Context) error {
	storage := models.Storages{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&storage)
	storage.ID = id

	if err := database.DB.Where("id= ?", storage.ID).Updates(&storage).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Storage Laptop",
		"storage": storage,
	})
}

func DeleteStorageByID(e echo.Context) error {
	var storage models.Storages
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&storage)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"storage": storage,
		"message": "Data Storage Berhasil Dihapus",
	})
}
