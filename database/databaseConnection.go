package database

import (
	"laptop_catalog/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	//hubungin ke database
	dsn := "root:@tcp(127.0.0.1:3306)/laptop_catalog?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:@tcp(host.docker.internal:3306)/laptop_catalog?charset=utf8mb4&parseTime=True&loc=Local"
	//kalau ada masalah dalam pemanggilan database, akan terjadi panic
	//panic:menghentikan proses yg sedg berlangsung, agar tidak terjadi infinite loop
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	initMigration()
}

var (
	DB *gorm.DB
)

func initMigration() {
	DB.AutoMigrate(&models.Admins{},
		&models.Descs{},
		&models.Types{},
		&models.Memories{},
		&models.Storages{},
		&models.Processors{},
		&models.Designs{},
		&models.Cameras{},
		&models.Batteries{},
		&models.Displays{},
		&models.Features{},
		&models.Connectors{},
		&models.Merks{})
}
