package models

type Customer struct {
	Id     int  `json:"id" binding:"required"`
	Drawed bool `json:"drawed" binding:"required"`
}
