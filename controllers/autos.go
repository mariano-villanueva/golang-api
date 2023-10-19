package controllers

import (
	db "api/database"
	"api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAutos(c *gin.Context) {
	var autos []models.Auto

	autos, err := db.GetAutos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error obteniendo elementos en base de datos.",
		})
		return
	}

	c.JSON(http.StatusOK, autos)
}

func PostAuto(c *gin.Context) {
	var nuevoAuto models.Auto

	if err := c.BindJSON(&nuevoAuto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
			"error":   err.Error(),
		})
		panic(err)
	}
	nuevoAuto, err := db.PostAuto(nuevoAuto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error insertando elemento en base de datos.",
		})
		panic(err)
	}
	c.JSON(http.StatusCreated, nuevoAuto)
}

func GetAutosPorCodigo(c *gin.Context) {
	codigo := c.Param("codigo")
	auto := db.GetAutoPorCodigo(codigo)
	c.JSON(http.StatusOK, auto)

}

func UpdateAuto(c *gin.Context) {
	var autoVacio models.Auto
	var nuevosDatos models.Auto
	nuevosDatos.Codigo = c.Param("codigo")
	err := c.BindJSON(&nuevosDatos)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
			"error":   err.Error(),
		})
		panic(err)
	}
	autoActualizado, err := db.UpdateAuto(nuevosDatos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error actualizando elemento en base de datos.",
		})
		panic(err)
	}
	if autoActualizado == autoVacio {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Error actualizando elemento inexistente.",
		})
		return
	}
	c.JSON(http.StatusOK, autoActualizado)
}

func DeleteAuto(c *gin.Context) {
	codigo := c.Param("codigo")
	db.DeleteAuto(codigo)
	c.JSON(http.StatusOK, gin.H{
		"message": "Recurso eliminado con éxito.",
	})
}
