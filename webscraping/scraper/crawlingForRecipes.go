package scraper

import (
	"fmt"
	"log"
)

func CrawlForWebPageRecipes(recipe string) {
	page, pw, browser := Initialize()
	urlRecipesSite := "https://theviewfromgreatisland.com/#search/q="
	siteQuery := urlRecipesSite + recipe
	urlRecipesSet := urlRecipeSet{}

	if _, err := page.Goto(siteQuery); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.WaitForSelector("div#card.vertical")

	entries, err := page.QuerySelectorAll("div#card.vertical")
	AssertErrorToNil("error getting selector: ", err)

	p := len(entries) // check if it is not an empty list
	fmt.Printf("Number of entries found: %v\n", p)

	for i, entry := range entries {
		recipeNameElement, err := entry.QuerySelector("a#name > span")
		AssertErrorToNil("could not get recipe element: ", err)

		recipeName, err := recipeNameElement.TextContent()
		AssertErrorToNil("could not get recipe name: ", err)

		linkElement, err := entry.QuerySelector("a#name")
		AssertErrorToNil("could not get link element: ", err)

		link, err := linkElement.GetAttribute("href")
		AssertErrorToNil("could not get link: ", err)
		urlRecipesSet.AddElement(link)

		fmt.Printf("%d:, recipe: %s, url: %s\n", i+1, recipeName, link)

	}
	urlsMap := urlsMap{}
	urlsMap.saveToMap(urlRecipesSet, recipe)
	jsonFormat := urlsMap.convertToJSON()
	saveJsonToFile(jsonFormat, recipe)
	//fmt.Println(urlsMap)
	Close(pw, browser)

}
