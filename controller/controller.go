package controller

import (
	"api/database"
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPersons(c *gin.Context) {
	var persons []models.Person //criando um novo array person que será do type Person
	database.DB.Find(&persons)
	c.JSON(200, persons) //mostra todas as persons do array
}

func GetByIdPerson(c *gin.Context) {
	var person models.Person    //criando uma nova var person que será do type Person
	id := c.Params.ByName("id") //pega toda a url vinda do usuário na função com c -> c *gin.Context é todo o retorno do usuário

	database.DB.First(&person, id) //ao encontrar o primeiro id que é igual ao digitado na url, retorna ele na hora -> .DB vai acessar todas as informações dentro da tabela do db

	if person.ID == 0 { //caso o aluno não exista -> colocar sempre depois da validação esse if e antes da mensagem que mostra o dado vindo do bd para que ele não apareça junto dessa mensagem
		c.JSON(http.StatusNotFound, "Not Found")
		return //sempre colocar o return para que não seja executado oque vem depois desse if
	}
	c.JSON(http.StatusOK, person) //mostra essa person encontrada com o id igual ao digitado na url
}

func SurchByCpf(c *gin.Context) {
	var person models.Person
	cpf := c.Param("cpf") //forma menor de buscar na url uma palavra chave

	database.DB.Where(&models.Person{Cpf: cpf}).First(&person) //faz uma busca no bd, deve passar a struct toda onde deve ser buscado e o valor logo em seguida

	if person.ID == 0 { //caso o aluno não exista -> colocar sempre depois da validação esse if e antes da mensagem que mostra o dado vindo do bd para que ele não apareça junto dessa mensagem
		c.JSON(http.StatusNotFound, "Not Found")
		return //sempre colocar o return para que não seja executado oque vem depois desse if
	}
	c.JSON(http.StatusOK, person) //mostra a person achada pelo cpf
}

func PostPerson(c *gin.Context) { //manda a informação para o banco
	var person models.Person

	if err := c.ShouldBindJSON(&person); err != nil { //caso um erro ocorra no endereço de memória que está salvo o &person
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Create(&person)   //salva os dados do endereço de memória que está salvo o &person
	c.JSON(http.StatusOK, person) //caso de tudo certo retorna o valor de person inserido
}

func PatchPerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id") //params é toda a url

	database.DB.First(&person, id) //acha a person no bd com o mesmo id digitado na url

	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	database.DB.Model(&person).UpdateColumns(person) //alterando o valor no bd
	c.JSON(http.StatusOK, "Alterado com sucesso")
}

func DeletePerson(c *gin.Context) {
	var person models.Person
	id := c.Params.ByName("id")

	database.DB.Delete(&person, id) //sempre que for realizar alguma alteração nas informações do banco usar o .DB
	c.JSON(http.StatusOK, "Deletada com sucesso")
}
