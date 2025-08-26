package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Marca": "fiat", "Modelo": "argo", "Ano": "2015", "Fipe": 45000.7, "Combustivel": "gasolina",
		})

	})
	r.Run(":8077")
}
