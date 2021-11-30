package models

type Cameras struct {
	ID                int    `json:"id" form:"id"`
	Webcam_Resolution string `json:"webcam_resolution" form:"webcam_resolution"`
}
