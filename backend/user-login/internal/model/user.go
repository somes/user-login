package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	// Name      string `gorm:"type:varchar(20); not null"`
	// Telephone string `gorm:"type:varchar(11); not null; unique"`
	// Password  string `gorm:"size:255; not null"`
	UserId        string `gorm:"type:varchar(36); not null; unique"`
	Username      string `gorm:"type:varchar(20); not null; unique"`
	Password      string `gorm:"type:varchar(20); not null"`
	RoleId        string `gorm:"type:int(2); not null"`
	Email         string `gorm:"type:varchar(32)"`
	Phone         string `gorm:"type:int(11)"`
	Status        string `gorm:"type:int(2); not null"`
	CertificateId string `gorm:"type:varchar(20)"`
}
