package router

import (
	"api/controller"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/", controller.GetPersons)
	r.GET("/:id", controller.GetByIdPerson)
	r.GET("/cpf/:cpf", controller.SurchByCpf)
	r.POST("/AddPerson", controller.PostPerson)
	r.PATCH("/:id", controller.PatchPerson)
	r.DELETE("/:id", controller.DeletePerson)
	r.Run() //não estava executando o servidor por conta que estava sem essa linha
}
