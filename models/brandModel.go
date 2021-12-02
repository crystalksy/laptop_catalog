package models

type Brands struct {
	ID    int    `json:"id" form:"id"`
	Brand string `json:"brand" form:"brand"`
}
