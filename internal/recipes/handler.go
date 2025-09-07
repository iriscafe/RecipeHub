package recipes

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"RecipeHub/internal/platform"
)

type RecipeHandler struct {
	service *RecipeService
}

func NewRecipeHandler(service *RecipeService) *RecipeHandler {
	return &RecipeHandler{
		service: service,
	}
}

func (h *RecipeHandler) GetAllRecipes(c *gin.Context) {
	platform.LogInfo("Buscando todas as receitas")
	recipes, err := h.service.GetAllRecipes()
	if err != nil {
		platform.LogError("Erro ao buscar receitas: " + err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	platform.LogInfo("Receitas encontradas: " + fmt.Sprintf("%d", len(recipes)))
	c.IndentedJSON(http.StatusOK, recipes)
}

func (h *RecipeHandler) GetRecipeById(c *gin.Context) {
	id := c.Param("id")
	recipe, err := h.service.GetRecipeById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	if recipe == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Recipe not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, recipe)
}

func (h *RecipeHandler) CreateRecipe(c *gin.Context) {
	var newRecipe Recipe
	if err := c.BindJSON(&newRecipe); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	
	err := h.service.CreateRecipe(newRecipe)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	
	c.IndentedJSON(http.StatusCreated, newRecipe)
}

func (h *RecipeHandler) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")
	var updatedRecipe Recipe
	if err := c.BindJSON(&updatedRecipe); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
		return
	}
	
	updatedRecipe.ID = id
	err := h.service.UpdateRecipe(updatedRecipe)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, updatedRecipe)
}

func (h *RecipeHandler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteRecipe(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}

// HealthCheck - Liveness probe
func (h *RecipeHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
		"service": "recipehub",
		"timestamp": platform.GetCurrentTime(),
	})
}

// ReadinessCheck - Readiness probe
func (h *RecipeHandler) ReadinessCheck(c *gin.Context) {
	// Verifica se o banco de dados está acessível
	if platform.DB == nil {
		platform.LogError("Database connection is nil")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": "not ready",
			"reason": "database not connected",
			"timestamp": platform.GetCurrentTime(),
		})
		return
	}

	platform.LogInfo("Readiness check passed")
	c.JSON(http.StatusOK, gin.H{
		"status": "ready",
		"service": "recipehub",
		"timestamp": platform.GetCurrentTime(),
	})
}