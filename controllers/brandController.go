package controllers

import (
	"laptop_catalog/database"
	"laptop_catalog/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBrands(c echo.Context) error {
	var brands []models.Brands

	if err := database.DB.Find(&brands).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Semua Merk Laptop",
		"brands":  brands,
	})
}

func GetBrandByID(c echo.Context) error {
	var brand models.Brands
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Where("id= ?", id).Find(&brand).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menampilkan Merk Laptop ",
		"brand":   brand,
	})
}

func CreateBrand(e echo.Context) error {
	brand := models.Brands{}
	e.Bind(&brand)

	if err := database.DB.Save(&brand).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Menambahkan Merk Laptop",
		"brand":   brand,
	})
}

func UpdateBrandByID(e echo.Context) error {
	brand := models.Brands{}
	id, _ := strconv.Atoi(e.Param("id"))
	e.Bind(&brand)
	brand.ID = id

	if err := database.DB.Where("id= ?", brand.ID).Updates(&brand).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Berhasil Mengubah Data Merk Laptop",
		"brand":   brand,
	})
}

//Fungsi hapus data merk
func DeleteBrandByID(e echo.Context) error {
	var brand models.Brands
	id, _ := strconv.Atoi(e.Param("id"))
	database.DB.Where("id = ?", id).Delete(&brand)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"brand":   brand,
		"message": "Data Merk Laptop Berhasil Dihapus",
	})
}
