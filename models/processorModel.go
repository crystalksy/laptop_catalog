package models

type Processors struct {
	ID               int    `json:"id" form:"id"`
	Operation_System string `json:"operation_system" form:"operation_system"`
	CPU              string `json:"cpu" form:"cpu"`
}
