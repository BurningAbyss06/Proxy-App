package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Nombre      string `gorm:"not null"`
	Apellido    string `gorm:"not null"`
	Contraseña  string `gorm:"not null"`
	Rol         string `gorm:"not null"`
	Foto_Perfil []byte `gorm:"type:bytea"`
	Decks       []Deck
	Matchs      []Match
}
