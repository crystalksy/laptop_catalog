package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

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

func GetConnectorByID(c echo.Context) error {
	var connector models.Connectors
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&connector).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "Berhasil Menampilkan Connector Laptop",
		"connector": connector,
	})
}

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

func DeleteConnectorByID(e echo.Context) error {
	var connector models.Connectors
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&connector)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"connector": connector,
		"message":   "Data Connector Berhasil Dihapus",
	})
}
