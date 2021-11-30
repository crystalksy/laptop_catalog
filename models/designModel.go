package models

type Designs struct {
	ID            int    `json:"id" form:"id"`
	Body_Material string `json:"body_material" form:"body_material"`
	Thickness     string `json:"thickness" form:"thickness"`
	Weight        string `json:"weight" form:"weight"`
}
