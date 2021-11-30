package models

type Batteries struct {
	ID            int    `json:"id" form:"id"`
	Daya          string `json:"daya" form:"daya"`
	Fast_Charging string `json:"fast_charging" form:"fast_charging"`
}
