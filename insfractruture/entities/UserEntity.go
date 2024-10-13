package entities

import "time"

type User struct {
	Id           uint       `gorm:"primary_key:auto_increment"`
	FirstName    string     `gorm:"type:varchar(150);not null"`
	FirstSurName string     `gorm:"type:varchar(150);not null" `
	SeconSurName string     `gorm:"type:varchar(150)"`
	Email        string     `gorm:"type:varchar(150);unique_index;not null"`
	Address      string     `gorm:"type:varchar(150)"`
	Phone        string     `gorm:"type:varchar(150);"`
	ZipCode      string     `gorm:"type:varchar(150)"`
	StateId      uint       `gorm:"type:integer"`
	RolId        uint       `gorm:"type:integer"`
	Active       bool       `gorm:"type:boolean"  json:"active"`
	CreatedAt    time.Time  `gorm:"<-:created_at"`
	UpdatedAt    *time.Time `gorm:"type:TIMESTAMP(6)"`
	Password     string     `gorm:"type:varchar(255);not null"`
}
