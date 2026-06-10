package howtocook

// Recipe represents a cooking recipe
type Recipe struct {
	ID               string       `json:"id"`
	Name             string       `json:"name"`
	Description      string       `json:"description"`
	SourcePath       string       `json:"source_path"`
	ImagePath        *string      `json:"image_path"`
	Category         string       `json:"category"`
	Difficulty       int          `json:"difficulty"`
	Tags             []string     `json:"tags"`
	Servings         int          `json:"servings"`
	Ingredients      []Ingredient `json:"ingredients"`
	Steps            []Step       `json:"steps"`
	PrepTimeMinutes  *int         `json:"prep_time_minutes"`
	CookTimeMinutes  *int         `json:"cook_time_minutes"`
	TotalTimeMinutes *int         `json:"total_time_minutes"`
	AdditionalNotes  []string     `json:"additional_notes"`
}

// Ingredient represents a recipe ingredient
type Ingredient struct {
	Name         string   `json:"name"`
	Quantity     *float64 `json:"quantity"`
	Unit         *string  `json:"unit"`
	TextQuantity string   `json:"text_quantity"`
	Notes        string   `json:"notes"`
}

// Step represents a cooking step
type Step struct {
	Step        int    `json:"step"`
	Description string `json:"description"`
}

// SimpleRecipe represents a simplified recipe
type SimpleRecipe struct {
	ID          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Ingredients []SimpleIngredient `json:"ingredients"`
}

// SimpleIngredient represents a simplified ingredient
type SimpleIngredient struct {
	Name         string `json:"name"`
	TextQuantity string `json:"text_quantity"`
}

// NameOnlyRecipe represents a recipe with only name and description
type NameOnlyRecipe struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// MealPlan represents a weekly meal plan
type MealPlan struct {
	Weekdays    []DayPlan   `json:"weekdays"`
	Weekend     []DayPlan   `json:"weekend"`
	GroceryList GroceryList `json:"groceryList"`
}

// DayPlan represents a daily meal plan
type DayPlan struct {
	Day       string         `json:"day"`
	Breakfast []SimpleRecipe `json:"breakfast"`
	Lunch     []SimpleRecipe `json:"lunch"`
	Dinner    []SimpleRecipe `json:"dinner"`
}

// GroceryList represents a grocery shopping list
type GroceryList struct {
	Ingredients  []GroceryItem `json:"ingredients"`
	ShoppingPlan ShoppingPlan  `json:"shoppingPlan"`
}

// GroceryItem represents an item in the grocery list
type GroceryItem struct {
	Name          string   `json:"name"`
	TotalQuantity *float64 `json:"totalQuantity"`
	Unit          *string  `json:"unit"`
	RecipeCount   int      `json:"recipeCount"`
	Recipes       []string `json:"recipes"`
}

// ShoppingPlan represents categorized shopping items
type ShoppingPlan struct {
	Fresh  []string `json:"fresh"`
	Pantry []string `json:"pantry"`
	Spices []string `json:"spices"`
	Others []string `json:"others"`
}

// DishRecommendation represents dish recommendations
type DishRecommendation struct {
	PeopleCount        int            `json:"peopleCount"`
	MeatDishCount      int            `json:"meatDishCount"`
	VegetableDishCount int            `json:"vegetableDishCount"`
	Dishes             []SimpleRecipe `json:"dishes"`
	Message            string         `json:"message"`
}
