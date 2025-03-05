package models

type Personalidade struct {
	Id       uint    `json:"id" gorm:"primaryKey"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Personalidade
