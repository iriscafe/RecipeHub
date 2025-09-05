package main

import (
	"github.com/gin-gonic/gin"
	"RecipeHub/internal/recipes"
	"RecipeHub/internal/platform"
)

func main() {

	platform.InitLogger()
	platform.InitConfig()
	platform.InitDB()
	

	platform.DB.AutoMigrate(&recipes.Recipe{})
	

	platform.LogInfo("Iniciando RecipeHub...")

	repository := recipes.NewRecipeRepository()
	service := recipes.NewRecipeService(repository)
	handler := recipes.NewRecipeHandler(service)

	router := gin.Default()
	router.GET("/recipes", handler.GetAllRecipes)
	router.GET("/recipes/:id", handler.GetRecipeById)
	router.POST("/recipes", handler.CreateRecipe)
	router.PUT("/recipes/:id", handler.UpdateRecipe)
	router.DELETE("/recipes/:id", handler.DeleteRecipe)
	
	port := ":" + platform.AppConfig.Port
	platform.LogInfo("Servidor rodando na porta " + platform.AppConfig.Port)
	platform.LogInfo("Ambiente: " + platform.AppConfig.Environment)
	platform.LogInfo("Nível de log: " + platform.AppConfig.LogLevel)
	
	router.Run(port)
}