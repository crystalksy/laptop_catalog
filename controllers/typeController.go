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
func GetTypes(c echo.Context) error {
	var types []models.Types

	if err := database.DB.Find(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Tipe Laptop",
		"types":   types,
	})
}

//Fungsi get admin by ID
func GetTypeByID(c echo.Context) error {
	var types models.Types
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Tipe Laptop ",
		"types":   types,
	})
}

//fungsi create new admins
func CreateType(e echo.Context) error {
	types := models.Types{}
	e.Bind(&types)

	if err := database.DB.Save(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Tipe Laptop",
		"type":    types,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateTypeByID(e echo.Context) error {
	types := models.Types{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&types)
	types.ID = id

	if err := database.DB.Where("id= ?", types.ID).Updates(&types).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Type Laptop",
		"type":    types,
	})
}

//Fungsi hapus data merk
func DeleteTypeByID(e echo.Context) error {
	var types models.Types
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&types)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"tipe":    types,
		"message": "Data Type Berhasil Dihapus",
	})
}
