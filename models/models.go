package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model         //ja adiciona ao criar o bd um id e o dia que foi criado, ultimo update e dia que foi excluido caso tenha sido
	Name       string  `json:"name"`
	Birthday   string  `json:"birthday"`
	Age        int     `json:"age"`
	Cpf        string  `json:"cpf"`
	Income     float32 `json:"income"`
	Height     float32 `json:"height"`
	Active     bool    `json:"active"`
	Obs        string  `json:"obs"`
}
