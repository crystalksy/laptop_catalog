package models

type Displays struct {
	ID                int    `json:"id" form:"id"`
	Screen_Dimension  int    `json:"screen_dimension" form:"screen_dimension"`
	Screen_Resolution int    `json:"screen_resolution" form:"screen_resolution"`
	Screen_Protector  string `json:"screen_Protector" form:"screen_Protector"`
	Screen_Tech       string `json:"screen_Tech" form:"screen_Tech"`
}
