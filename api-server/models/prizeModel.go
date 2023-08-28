package models

type Prize struct {
	Category    string  `json:"category" binding:"required"`
	Probability float64 `json:"probability" binding:"required"`
	Totalquota  int     `json:"totalquota" binding:"required"`
	Dailyquota  int     `json:"dailyquota" binding:"required"`
}
