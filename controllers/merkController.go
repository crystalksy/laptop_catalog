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
func GetMerks(c echo.Context) error {
	var merks []models.Merks

	if err := database.DB.Find(&merks).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Merk Laptop",
		"merks":   merks,
	})
}

//Fungsi get admin by ID
func GetMerkByID(c echo.Context) error {
	var merk models.Merks
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&merk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Merk Laptop ",
		"merk":    merk,
	})
}

//fungsi create new admins
func CreateMerk(e echo.Context) error {
	merk := models.Merks{}
	e.Bind(&merk)

	if err := database.DB.Save(&merk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Merk Laptop",
		"merk":    merk,
	})
}

//Fungsi Update Tabel Merk HP
func UpdateMerkByID(e echo.Context) error {
	merk := models.Merks{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&merk)
	merk.ID = id

	if err := database.DB.Where("id= ?", merk.ID).Updates(&merk).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Merk Laptop",
		"merk":    merk,
	})
}

//Fungsi hapus data merk
func DeleteMerkByID(e echo.Context) error {
	var merk models.Merks
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&merk)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"merk":    merk,
		"message": "Data Merk Laptop Berhasil Dihapus",
	})
}
