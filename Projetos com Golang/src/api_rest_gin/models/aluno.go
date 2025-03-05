package models

import "gorm.io/gorm"

type Aluno struct {
	gorm.Model
	// Informamos como vamos vincular esse campo com a resposta json
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}