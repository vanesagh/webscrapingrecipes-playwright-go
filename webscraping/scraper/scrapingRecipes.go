package scraper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

type Ingredients struct {
	Ingredients []Ingredient `json:"ingredients"`
}

func (i *Ingredients) AddIngredient(ingredient *Ingredient) {
	i.Ingredients = append(i.Ingredients, *ingredient)

}

type Ingredient struct {
	Name   string `json:"item"`
	Amount string `json:"amount"`
	Unit   string `json:"unit"`
}

func GetAmountQuantityName(selectorPath string, ingredient playwright.ElementHandle, message string) string {
	defer handlePanic()
	amount_selector, err := ingredient.QuerySelector(selectorPath)
	AssertErrorToNil("could not get "+message+" %v", err)
	if amount_selector == nil {
		panic(message)
	} else {
		amount, err := amount_selector.TextContent()
		AssertErrorToNil("could not get text %v", err)
		//fmt.Printf(" %s\n", amount)
		return amount

	}
}

func ScrapForRecipe(urlRecipe string) {
	page, pw, browser := Initialize()
	if _, err := page.Goto(urlRecipe); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	page.WaitForSelector("div.wprm-recipe-ingredient-group")

	recipeNameSelector, err := page.QuerySelector("h2.wprm-recipe-name")
	AssertErrorToNil("error getting selector h2: ", err)
	recipeName, err := recipeNameSelector.TextContent()
	AssertErrorToNil("could not get recipe name text", err)
	fmt.Print(recipeName)

	entries, err := page.QuerySelectorAll("div.wprm-recipe-ingredient-group")
	AssertErrorToNil("error getting selector(s): ", err)

	p := len(entries) // check if it is not an empty list
	fmt.Printf("Number of entries found: %v\n", p)
	wholeIngredients := make(map[string][]Ingredient)
	for _, entry := range entries {

		process_selector, err := entry.QuerySelector("h4.wprm-recipe-group-name")
		AssertErrorToNil("error getting process selector", err)
		process, err := process_selector.TextContent()
		AssertErrorToNil("could not get process text", err)

		fmt.Printf("process: %s ,", process)
		//arrayIngredients := Ingredients{}

		ingredients, err := entry.QuerySelectorAll("ul.wprm-recipe-ingredients > li")
		AssertErrorToNil("error getting selectors %v", err)
		p = len(ingredients)
		fmt.Printf("Number of ingredients: %d\n", p)

		for _, ingredient := range ingredients {

			amount := GetAmountQuantityName("span.wprm-recipe-ingredient-amount", ingredient, "no amount found")
			unit := GetAmountQuantityName("span.wprm-recipe-ingredient-unit", ingredient, "no unit found")
			name := GetAmountQuantityName("span.wprm-recipe-ingredient-name", ingredient, "no name found")

			ingred := &Ingredient{
				Amount: amount,
				Unit:   unit,
				Name:   name,
			}
			//arrayIngredients.AddIngredient(ingred)
			wholeIngredients[process] = append(wholeIngredients[process], *ingred)

		}
		//wholeIngredients[process] = arr
	}
	//fmt.Printf("%+v\n", wholeIngredients)
	jsonFormat, err := json.Marshal(wholeIngredients)
	AssertErrorToNil("error converting to json", err)
	fmt.Println(string(jsonFormat))
	saveJsonToFile(jsonFormat, recipeName)
	Close(pw, browser)

}

func AssertErrorToNil(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}

}

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("Recover of: ", a)
	}
}
