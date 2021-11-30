package controllers

//import packages
import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data displays
func GetDisplays(c echo.Context) error {
	var display []models.Displays

	if err := database.DB.Find(&display).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Display Laptop",
		"display": display,
	})
}

//Fungsi get displaysby ID
func GetDisplayByID(c echo.Context) error {
	var display models.Displays
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&display).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Displays Laptop ",
		"display": display,
	})
}

//fungsi create new displays
func CreateDisplay(e echo.Context) error {
	display := models.Displays{}
	e.Bind(&display)

	if err := database.DB.Save(&display).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "success menambahkan data Display Laptop",
		"display": display,
	})
}

//Fungsi Update Tabel Displays
func UpdateDisplayByID(e echo.Context) error {
	display := models.Displays{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&display)
	display.ID = id

	if err := database.DB.Where("id= ?", display.ID).Updates(&display).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Display",
		"display": display,
	})
}

//Fungsi hapus data displays
func DeleteDisplayByID(e echo.Context) error {
	var display models.Displays
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&display)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"display": display,
		"message": "Data Berhasil Dihapus",
	})
}
