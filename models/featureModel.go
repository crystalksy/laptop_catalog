package models

type Features struct {
	ID               int    `json:"id" form:"id"`
	Backlit_Keyboard string `json:"backlit_keyboard" form:"backlit_keyboard"`
	Touchscreen      string `json:"touchscreen" form:"touchscreen"`
	Fingerprint      string `json:"fingerprint" form:"fingerprint"`
}
