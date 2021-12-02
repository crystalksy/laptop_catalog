package test

import (
	"fmt"
	"laptop_catalog/database"
	"laptop_catalog/models"
	"laptop_catalog/route"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/labstack/echo/v4"
)

var (
	echoHandler *echo.Echo
	server      *httptest.Server
)

func init() {
	database.InitDB("laptop_catalog_test")
	echoHandler = route.New()
	server = httptest.NewServer(echoHandler)

}

func TearUp() func() {
	database.InitDB("laptop_catalog_test")

	laptop := &models.Types{
		Price:        "Test1",
		Release_Year: "Test2",
	}
	if err := database.DB.Create(&laptop).Error; err != nil {
		fmt.Println("error")
	}
	laptop1 := &models.Brands{
		Brand: "Test3",
	}
	if err := database.DB.Create(&laptop1).Error; err != nil {
		fmt.Println("error")
	}
	return func() {

		database.InitDB("laptop_catalog_test")
		database.DB.Exec("TRUNCATE TABLE types;")
		database.DB.Exec("TRUNCATE TABLE brands;")
	}
}

func GetHTTPExpect(t *testing.T) *httpexpect.Expect {
	if server == nil {
		server = httptest.NewServer(echoHandler)
	}
	return NewHTTPExpect(t)
}

func NewHTTPExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: server.URL,
		Printers: []httpexpect.Printer{
			httpexpect.NewCurlPrinter(t),
			httpexpect.NewDebugPrinter(t, true),
		},
		Reporter: t,
	})
}
