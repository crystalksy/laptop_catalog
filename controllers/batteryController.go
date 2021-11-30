package controllers

//import packages
import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data admin
func GetBatteries(c echo.Context) error {
	var battery []models.Batteries

	if err := database.DB.Find(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Baterai Laptop",
		"battery": battery,
	})
}

//Fungsi get admin by ID
func GetBatteryByID(c echo.Context) error {
	var battery models.Batteries
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Baterai Laptop ",
		"battery": battery,
	})
}

//fungsi create new admins
func CreateBattery(e echo.Context) error {
	battery := models.Batteries{}
	e.Bind(&battery)

	if err := database.DB.Save(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Battery Laptop",
		"battery": battery,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateBatteryByID(e echo.Context) error {
	battery := models.Batteries{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&battery)
	battery.ID = id

	if err := database.DB.Where("id= ?", battery.ID).Updates(&battery).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Baterai",
		"battery": battery,
	})
}

//Fungsi hapus data merk
func DeleteBatteryByID(e echo.Context) error {
	var battery models.Batteries
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&battery)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"battery": battery,
		"message": "Data Baterai Berhasil Dihapus",
	})
}
