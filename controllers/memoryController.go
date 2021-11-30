package controllers

//import packages
import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data audio
func GetMemories(c echo.Context) error {
	var memory []models.Memories

	if err := database.DB.Find(&memory).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Data Memory Laptop",
		"memory":  memory,
	})
}

//Fungsi get audio by ID
func GetMemoryByID(c echo.Context) error {
	var memory models.Memories
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&memory).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Storage Laptop ",
		"memory":  memory,
	})
}

//fungsi create new audio
func CreateMemory(e echo.Context) error {
	memory := models.Memories{}
	e.Bind(&memory)

	if err := database.DB.Save(&memory).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Data Memory Laptop",
		"memory":  memory,
	})
}

//Fungsi Update Tabel Audio
func UpdateMemoryByID(e echo.Context) error {
	memory := models.Memories{}
	e.Bind(&memory)

	if err := database.DB.Where("id= ?", memory.ID).Updates(&memory).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Memory Laptop",
		"memory":  memory,
	})
}

//Fungsi hapus data audio
func DeleteMemoryByID(e echo.Context) error {
	var memory models.Memories
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&memory)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"memory":  memory,
		"message": "Data Storage Berhasil Dihapus",
	})
}
