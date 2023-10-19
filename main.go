package main

import (
	"api/controllers"
	db "api/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	router := gin.Default()

	rutasAutos := router.Group("/autos")
	{
		rutasAutos.GET("", controllers.GetAutos)
		rutasAutos.POST("", controllers.PostAuto)
		rutasAutos.GET("/:codigo", controllers.GetAutosPorCodigo)
		rutasAutos.PUT("/:codigo", controllers.UpdateAuto)
		rutasAutos.DELETE("/:codigo", controllers.DeleteAuto)
	}

	router.Run("localhost:3388")
}
