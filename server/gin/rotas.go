package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Carro struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Marca  string `json: "marca"`
	Modelo string `json: "modelo"`
	Km     int    `json: "km"`
}

var banco *gorm.DB

func main() {
	var err error

	banco, err = gorm.Open(sqlite.Open("carros.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro ao se conecar ao banco")
	}
	banco.AutoMigrate(&Carro{})
	r := gin.Default()
	//rotas
	r.GET("/carros", listarcarros)
	r.POST("/carros", criarCarros)
	r.POST("/varioscarros", criarvariosCarros)
	r.GET("/carros/:marca", buscarPorMarca)

	r.Run(":8099")
}

func listarcarros(c *gin.Context) {
	var carros []Carro
	banco.Find(&carros)
	c.JSON(http.StatusOK, carros)

}

func criarCarros(c *gin.Context) {
	var carro Carro
	if err := c.ShouldBindJSON(&carro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error})
		return
	}
	banco.Create(&carro)
	c.JSON(http.StatusCreated, carro)

}

func criarvariosCarros(c *gin.Context) {
	var carro []Carro
	if err := c.ShouldBindJSON(&carro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error})
		return
	}
	banco.Create(&carro)
	c.JSON(http.StatusCreated, carro)

}

func buscarPorMarca(c *gin.Context) {
	marca := c.Param("marca")
	var carros []Carro
	result := banco.Where("marca=?", marca).Find(&carros)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "carro nao encontrado"})
		return
	}
	if len(carros) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"mensagem": "Nenhum carro encontrado para essa marca"})
		return
	}
	c.JSON(http.StatusOK, carros)

}
