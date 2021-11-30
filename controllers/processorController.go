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
func GetProcessors(c echo.Context) error {
	var processor []models.Processors

	if err := database.DB.Find(&processor).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Semua Data Processor Laptop",
		"processor": processor,
	})
}

//Fungsi get audio by ID
func GetProcessorByID(c echo.Context) error {
	var processor models.Processors
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&processor).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Processor Laptop ",
		"processor": processor,
	})
}

//fungsi create new audio
func CreateProcessor(e echo.Context) error {
	processor := models.Processors{}
	e.Bind(&processor)

	if err := database.DB.Save(&processor).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menambahkan Data Processor Laptop",
		"processor": processor,
	})
}

//Fungsi Update Tabel Audio
func UpdateProcessorByID(e echo.Context) error {
	processor := models.Processors{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&processor)
	processor.ID = id

	if err := database.DB.Where("id= ?", processor.ID).Updates(&processor).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Mengubah Data Processor Laptop",
		"processor": processor,
	})
}

//Fungsi hapus data audio
func DeleteProcessorByID(e echo.Context) error {
	var processor models.Processors
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&processor)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"processor": processor,
		"message":   "Data Processor Berhasil Dihapus",
	})
}
