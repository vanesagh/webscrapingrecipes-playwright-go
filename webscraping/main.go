package main

import (
	scraper "webscraping/scraper"
)

func main() {
	//scrapForWebPageRecipes()
	scraper.ScrapForRecipe("https://theviewfromgreatisland.com/lemon-muffins-recipe/")

}
