package models

import "gorm.io/gorm"

type Match struct {
	gorm.Model

	UserID          uint `gorm:"index;foreignKey:UserID;references:users(ID)"`
	Deck_Jugador    Deck
	Deck_Oponente   Deck
	Record          string `gorm:"not null"`
	Dado            bool   `gorm:"default:true"`
	Matchup         string `gorm:"not null"`
	Observaciones   string
	Deck_JugadorID  uint `gorm:"foreignKey:Deck_JugadorID;references:deck(ID)"`
	Deck_OponenteID uint `gorm:"foreignKey:Deck_OponenteID;references:deck(ID)"`
}
