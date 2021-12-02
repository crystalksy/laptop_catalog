package main

import (
	"laptop_catalog/database"
	"laptop_catalog/route"
)

//MAIN FUNCTION
func main() {
	database.InitDB("laptop_catalog")
	e := route.New()
	e.Start(":8080")
}
