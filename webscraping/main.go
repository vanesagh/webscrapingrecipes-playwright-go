package main

import (
	scraper "webscraping/scraper"
)

func main() {
	scraper.CrawlForWebPageRecipes("muffins")
	//scraper.ScrapForRecipe("https://theviewfromgreatisland.com/lemon-muffins-recipe/")

}
