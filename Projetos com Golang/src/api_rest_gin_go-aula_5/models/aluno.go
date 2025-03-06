package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
