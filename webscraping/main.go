package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/playwright-community/playwright-go"
)

type Ingredients struct {
	//process     string
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

func handlePanic() {
	a := recover()

	if a != nil {
		fmt.Println("Recover of: ", a)
	}
}

func getAmountQuantityName(selectorPath string, ingredient playwright.ElementHandle, message string) string {
	defer handlePanic()
	amount_selector, err := ingredient.QuerySelector(selectorPath)
	assertErrorToNil("could not get "+message+" %v", err)
	if amount_selector == nil {
		panic(message)
	} else {
		amount, err := amount_selector.TextContent()
		assertErrorToNil("could not get text %v", err)
		//fmt.Printf(" %s\n", amount)
		return amount

	}
}

func assertErrorToNil(message string, err error) {
	if err != nil {
		log.Fatalf(message, err)
	}

}

func initialize() (playwright.Page, *playwright.Playwright, playwright.Browser) {
	pw, err := playwright.Run()
	assertErrorToNil("could not start playwright: ", err)

	browser, err := pw.Chromium.Launch()
	assertErrorToNil("could not launch browser: ", err)

	page, err := browser.NewPage()
	page.SetDefaultTimeout(10000) //needed more time. Otherwise got nothing
	assertErrorToNil("could not create page: ", err)

	return page, pw, browser

}

func close(pw *playwright.Playwright, browser playwright.Browser) {
	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

}

func scrapForRecipe() {
	page, pw, browser := initialize()
	if _, err := page.Goto("https://theviewfromgreatisland.com/lemon-muffins-recipe/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	page.WaitForSelector("div.wprm-recipe-ingredient-group")

	recipeNameSelector, err := page.QuerySelector("h2.wprm-recipe-name")
	assertErrorToNil("error getting selector h2: ", err)
	recipeName, err := recipeNameSelector.TextContent()
	assertErrorToNil("could not get recipe name text", err)
	fmt.Printf(recipeName)

	entries, err := page.QuerySelectorAll("div.wprm-recipe-ingredient-group")
	assertErrorToNil("error getting selector(s): ", err)

	p := len(entries) // check if it is not an empty list
	fmt.Printf("Number of entries found: %v\n", p)
	wholeIngredients := make(map[string]Ingredients)
	for _, entry := range entries {

		process_selector, err := entry.QuerySelector("h4.wprm-recipe-group-name")
		assertErrorToNil("error getting process selector", err)
		process, err := process_selector.TextContent()
		assertErrorToNil("could not get process text", err)

		fmt.Printf("process: %s ,", process)
		arrayIngredients := Ingredients{}

		ingredients, err := entry.QuerySelectorAll("ul.wprm-recipe-ingredients > li")
		assertErrorToNil("error getting selectors %v", err)
		p = len(ingredients)
		fmt.Printf("Number of ingredients: %d\n", p)

		for _, ingredient := range ingredients {

			amount := getAmountQuantityName("span.wprm-recipe-ingredient-amount", ingredient, "no amount found")
			unit := getAmountQuantityName("span.wprm-recipe-ingredient-unit", ingredient, "no unit found")
			name := getAmountQuantityName("span.wprm-recipe-ingredient-name", ingredient, "no name found")

			ingredient := &Ingredient{
				Amount: amount,
				Unit:   unit,
				Name:   name,
			}
			arrayIngredients.AddIngredient(ingredient)

		}
		wholeIngredients[process] = arrayIngredients
	}
	//fmt.Printf("%+v\n", wholeIngredients)
	jsonFormat, err := json.Marshal(wholeIngredients)
	assertErrorToNil("error converting to json", err)
	fmt.Println(string(jsonFormat))
	saveJsonToFile(jsonFormat, recipeName)
	close(pw, browser)

}

func saveJsonToFile(jsonFormat []byte, recipeName string) {
	fileName := recipeName + ".json"
	err := ioutil.WriteFile(fileName, jsonFormat, 0644)
	assertErrorToNil("error saving file", err)

}

func scrapForWebPageRecipes() {
	page, pw, browser := initialize()

	if _, err := page.Goto("https://theviewfromgreatisland.com/#search/q=muffins"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.WaitForSelector("div#card.vertical")
	entries, err := page.QuerySelectorAll("div#card.vertical")

	assertErrorToNil("error getting selector: ", err)

	p := len(entries) // check if it is not an empty list
	fmt.Printf("Number of entries found: %v\n", p)

	for i, entry := range entries {
		recipeNameElement, err := entry.QuerySelector("a#name > span")
		assertErrorToNil("could not get recipe element: ", err)

		recipeName, err := recipeNameElement.TextContent()
		assertErrorToNil("could not get recipe name: ", err)

		linkElement, err := entry.QuerySelector("a#name")
		assertErrorToNil("could not get link element: ", err)

		link, err := linkElement.GetAttribute("href")
		assertErrorToNil("could not get link: ", err)

		fmt.Printf("%d:, recipe: %s, url: %s\n", i+1, recipeName, link)

	}
	close(pw, browser)

}

func main() {
	//scrapForWebPageRecipes()
	scrapForRecipe()

}
