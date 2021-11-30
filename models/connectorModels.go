package models

type Connectors struct {
	ID             int    `json:"id" form:"id"`
	HDMI           int    `json:"hdmi" form:"hdmi"`
	USB_TypeC      string `json:"usb_typec" form:"usb_typec"`
	Headphone_Jack string `json:"headphone_jack" form:"headphone_jack"`
	VGA            string `json:"vga" form:"vga"`
	Display_Port   string `json:"display_port" form:"display_port"`
	Thunderbolt    string `json:"thunderbolt" form:"thunderbolt"`
	USB            string `json:"usb" form:"usb"`
	Card_Reader    string `json:"card_reader" form:"card_reader"`
	DVI            string `json:"dvi" form:"dvi"`
}
