package routers

import (
	"github.com/gin-gonic/gin"
	"recipe-api/controllers"
)

func GetRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/recipes", controllers.Recipes)

	return router
}
