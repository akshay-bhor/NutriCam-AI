package models

type Users struct {
	Id          uint    `gorm:"primaryKey;autoIncrement;unique;not null:true" json:"id"`
	User        string  `gorm:"unique;not null:true" json:"user"`
	Mail        string  `gorm:"unique,not null:true;size:50" json:"mail"`
	Gid         *string `gorm:"unique,not null:false;size:21" json:"gid"`
	Pass        *string `gorm:"size:60;not null:false" json:"pass"`
	Totp_secret string  `gorm:"not null:true" json:"totp_secret"`
	Name        string  `gorm:"not null:true" json:"name"`
	Surname     string  `gorm:"not null:true" json:"surname"`
	Status      string  `gorm:"not null:true" json:"status"`
	Type        string  `gorm:"not null:true;default:user" json:"type"`
}
