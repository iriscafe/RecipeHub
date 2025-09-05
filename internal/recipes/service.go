package recipes

type RecipeService struct {
	repository RecipeRepository
}

func NewRecipeService(repository RecipeRepository) *RecipeService {
	return &RecipeService{
		repository: repository,
	}
}

func (s *RecipeService) GetAllRecipes() ([]Recipe, error) {
	return s.repository.GetAllRecipes()
}

func (s *RecipeService) GetRecipeById(id string) (*Recipe, error) {
	return s.repository.GetRecipeById(id)
}

func (s *RecipeService) CreateRecipe(recipe Recipe) error {
	return s.repository.CreateRecipe(recipe)
}

func (s *RecipeService) UpdateRecipe(recipe Recipe) error {
	return s.repository.UpdateRecipe(recipe)
}

func (s *RecipeService) DeleteRecipe(id string) error {
	return s.repository.DeleteRecipe(id)
}
