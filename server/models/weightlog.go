package models

type WeightLog struct {
	UserId     uint    `gorm:"unique;not null:true" json:"user_id"`
	Weight     float64 `gorm:"not null:true;type:float(3, 1)" json:"weight"`
	Weight_lbs float64 `gorm:"not null:true;type:float(3, 1)" json:"weight_lbs"`
}
