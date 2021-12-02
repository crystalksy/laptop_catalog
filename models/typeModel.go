package models

type Types struct {
	ID           int    `json:"id" form:"id"`
	Id_Merk      int    `json:"id_merk" form:"id_merk"`
	Release_Year string `json:"release_year" form:"release_year"`
	Price        string `json:"price" form:"price"`
}
