package models

type Aluno struct {
	// Informamos como vamos vincular esse campo com a resposta json
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Alunos []Aluno