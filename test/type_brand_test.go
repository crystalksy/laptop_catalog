package test

import (
	"net/http"
	"testing"
)

func TestGetType(t *testing.T) {
	tearDown := TearUp()

	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/types").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("Berhasil Menampilkan Semua Tipe Laptop")

	result.Value("types").Array().NotEmpty()
}

func TestGetBrand(t *testing.T) {
	tearDown := TearUp()
	defer tearDown()

	e := GetHTTPExpect(t)

	result := e.GET("/brands").Expect().Status(http.StatusOK).JSON().Object()

	result.Value("message").String().Contains("Berhasil Menampilkan Semua Merk Laptop")
	result.Value("brands").Array().NotEmpty()
}
