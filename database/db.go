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

func ConectaBandoDeDados() {
	dsn := "host=localhost user=qaafopmjwydwrv password=f5a8ac091423ab38b08f97c02cff936ee0d7666c59df8eb579c7383685b3ed24 dbname=desq7v6u6e2f81 port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com bando de dados")
	}
}
