package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

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
	page.SetDefaultTimeout(6000) //needed more time. Otherwise got nothing
	assertErrorToNil("could not create page: ", err)

	return page, pw, browser

}

func scrapForWebPageRecipes() {
	page, pw, browser := initialize()

	if _, err := page.Goto("https://theviewfromgreatisland.com/#search/q=muffins"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	page.WaitForSelector("div#card.vertical")

	entries, err := page.QuerySelectorAll("div#card.vertical")

	if err != nil {
		log.Fatalf("error getting selector: %v", err)
	}

	p := len(entries) // check if it is not an empty list
	fmt.Printf("Number of entries found: %v\n", p)

	for i, entry := range entries {
		recipeNameElement, err := entry.QuerySelector("a#name > span")
		if err != nil {
			log.Fatalf("could not get recipe element: %v ", err)
		}

		recipeName, err := recipeNameElement.TextContent()
		if err != nil {
			log.Fatalf("could not get recipe name: %v ", err)
		}

		linkElement, err := entry.QuerySelector("a#name")
		if err != nil {
			log.Fatalf("could not get link element: %v", err)
		}

		link, err := linkElement.GetAttribute("href")
		if err != nil {
			log.Fatalf("could not get link: %v ", err)
		}
		fmt.Printf("%d:, recipe: %s, url: %s\n", i+1, recipeName, link)

	}
	if err = browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

}

func main() {
	scrapForWebPageRecipes()

}
