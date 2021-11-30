package models

type Storages struct {
	ID           int    `json:"id" form:"id"`
	Storage_Type string `json:"storage_type" form:"storage_type"`
	CPU          string `json:"cpu" form:"cpu"`
}
