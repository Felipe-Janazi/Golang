package models

import (
	"gorm.io/gorm"
	"gopkg.in/validator.v2"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero"`
	RG   string `json:"rg" validate:"len=9"`
	CPF  string `json:"cpf" validate:"len=11"  regexp=^[0-9]*$`
}

func ValidaDadosDeAluno(aluno *Aluno) error {

	if err := validator.Validate(aluno); err != nil{
		return err
	}

	return nil
}