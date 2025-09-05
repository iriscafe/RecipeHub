package recipes

import (
	"time"
	"RecipeHub/internal/platform"
	"gorm.io/gorm"
)

type Recipe struct {
	ID string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Description string `json:"description"`
	Instructions string `json:"instructions"`
	Ingredients string `json:"ingredients"`
	Category string `json:"category"`
	Subcategory string `json:"subcategory"`
	Tags string `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RecipeRepository interface {
	GetAllRecipes() ([]Recipe, error)
	GetRecipeById(id string) (*Recipe, error)
	CreateRecipe(recipe Recipe) error
	UpdateRecipe(recipe Recipe) error
	DeleteRecipe(id string) error
}

type recipeRepository struct {
	db *gorm.DB
}

func NewRecipeRepository() RecipeRepository {
	return &recipeRepository{
		db: platform.DB,
	}
}

func (r *recipeRepository) GetAllRecipes() ([]Recipe, error) {
	var recipes []Recipe
	result := r.db.Find(&recipes)
	return recipes, result.Error
}

func (r *recipeRepository) GetRecipeById(id string) (*Recipe, error) {
	var recipe Recipe
	result := r.db.First(&recipe, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &recipe, nil
}

func (r *recipeRepository) CreateRecipe(recipe Recipe) error {
	// Definir timestamps
	now := time.Now()
	recipe.CreatedAt = now
	recipe.UpdatedAt = now
	
	result := r.db.Create(&recipe)
	return result.Error
}

func (r *recipeRepository) UpdateRecipe(recipe Recipe) error {
	// Atualizar timestamp
	recipe.UpdatedAt = time.Now()
	
	result := r.db.Save(&recipe)
	return result.Error
}

func (r *recipeRepository) DeleteRecipe(id string) error {
	result := r.db.Delete(&Recipe{}, "id = ?", id)
	return result.Error
}
