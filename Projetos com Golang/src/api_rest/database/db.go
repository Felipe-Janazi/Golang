package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=UTC"
	DB, err = gorm.Open(postgres.Open(stringDeConexao), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar com o Banco de Dados")
	}
}
