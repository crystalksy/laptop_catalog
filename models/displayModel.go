package models

type Displays struct {
	ID                int    `json:"id" form:"id"`
	Screen_Dimension  string `json:"screen_dimension" form:"screen_dimension"`
	Screen_Resolution string `json:"screen_resolution" form:"screen_resolution"`
	Screen_Tech       string `json:"screen_Tech" form:"screen_Tech"`
}
