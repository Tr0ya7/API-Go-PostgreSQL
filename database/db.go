package database

import (
	"api/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConect() {
	conector := "host=localhost user=postgres password=masterkey dbname=teste port=5432 sslmode=disable" //informações de coneção
	DB, err = gorm.Open(postgres.Open(conector))                                                         //se conecta ao banco

	if err != nil { //caso ocorra qualquer erro
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Person{}) //vai criar uma tabela de pessoas no postgres com as informações nela
}
