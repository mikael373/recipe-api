package controllers

import (
	"github.com/gin-gonic/gin"
	"recipe-api/services"
	"sort"
	"strings"
)

func Recipes(c *gin.Context) {
	ingredientsString := c.Query("i")
	if ingredientsString == "" {
		c.JSON(400, gin.H{
			"erro": "nenhum ingrediente informado",
		})
		return
	}
	ingredients := strings.Split(ingredientsString, ",")
	sort.Strings(ingredients)
	recipes, err := services.GetRecipes(ingredients)
	if err != nil {
		c.JSON(500, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(200, recipes)
}
