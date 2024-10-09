package models

import "gorm.io/gorm"

type Deck struct {
	gorm.Model
	Apodo string `gorm:"not null"`
	/*MatchupID uint*/
	UserID uint `gorm:"index;foreignKey:UserID;references:users(ID)"`
}
