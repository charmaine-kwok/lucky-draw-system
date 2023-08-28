package models

type Mobile struct {
	Customerid int    `json:"-"`
	Mobile     string `json:"mobile" binding:"required" example:"98765432"`
}
