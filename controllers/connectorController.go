package controllers

//import packages
import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

//func memanggil seluruh data perfoms
func GetConnectors(c echo.Context) error {
	var connector []models.Connectors

	if err := database.DB.Find(&connector).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Semua Data Connector Laptop",
		"connector": connector,
	})
}

//Fungsi get perfoms by ID
func GetConnectorByID(c echo.Context) error {
	var connector models.Connectors
	id, _ := strconv.Atoi(c.Param("id"))
	//config.DB.Where("id = ?", id).Delete(&admin)

	if err := database.DB.Where("id= ?", id).Find(&connector).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Connector Laptop",
		"connector": connector,
	})
}

//fungsi create new perfoms
func CreateConnector(e echo.Context) error {
	connector := models.Connectors{}
	e.Bind(&connector)

	if err := database.DB.Save(&connector).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success menambahkan Data Connector",
		"connector": connector,
	})
}

//Fungsi Update Tabel Perfoms
func UpdateConnectorByID(e echo.Context) error {
	connector := models.Connectors{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&connector)
	connector.ID = id

	if err := database.DB.Where("id= ?", connector.ID).Updates(&connector).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Mengubah Data Connector",
		"connector": connector,
	})
}

//Fungsi hapus data perfoms
func DeleteConnectorByID(e echo.Context) error {
	var connector models.Connectors
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&connector)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"connector": connector,
		"message":   "Data Connector Berhasil Dihapus",
	})
}
