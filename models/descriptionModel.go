package models

type Descs struct {
	ID           int `json:"id" form:"id"`
	Id_Type      int `json:"id_type" form:"id_type"`
	Id_Design    int `json:"id_design" form:"id_design"`
	Id_Camera    int `json:"id_camera" form:"id_camera"`
	Id_Battery   int `json:"id_battery" form:"id_battery"`
	Id_Connector int `json:"id_connector" form:"id_connector"`
	Id_Display   int `json:"id_display" form:"id_display"`
	Id_Feature   int `json:"id_feature" form:"id_feature"`
	Id_Memory    int `json:"id_memory" form:"id_memory"`
	Id_Merk      int `json:"id_merk" form:"id_merk"`
	Id_Processor int `json:"id_processor" form:"id_processor"`
	Id_Storage   int `json:"id_storage" form:"id_storage"`
}
