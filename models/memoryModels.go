package models

type Memories struct {
	ID       int    `json:"id" form:"id"`
	RAM      string `json:"ram" form:"ram"`
	Type_RAM string `json:"type_ram" form:"type_ram"`
}
