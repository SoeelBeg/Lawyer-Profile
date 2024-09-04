package models

import "gorm.io/gorm"

type Lawyer struct {
	ID           int64  `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Specialty    string `json:"specialty"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Address      string `json:"address"`
}

func MigrateLawyer(db *gorm.DB) {
	db.AutoMigrate(&Lawyer{})
}
