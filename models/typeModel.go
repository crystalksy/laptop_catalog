package models

type Types struct {
	ID           int    `json:"id" form:"id"`
	Id_Merk      int    `json:"id_merk" form:"id_merk"`
	Type         string `json:"type" form:"type"`
	Release_Year int    `json:"release_year" form:"release_year"`
	Price        int    `json:"price" form:"price"`
}
