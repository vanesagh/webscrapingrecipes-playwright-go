package scraper

import (
	"fmt"
	"testing"
)

func TestScrapeRecipe(t *testing.T) {
	t.Run("handle wrong url", func(t *testing.T) {
		urlRecipe := "https://www.google.com"
		err := ScrapeForRecipe(urlRecipe)

		fmt.Println(err)

	})

	t.Run("handle recipes with one process", func(t *testing.T) {
		ScrapeForRecipe("https://theviewfromgreatisland.com/triple-berry-muffins-recipe/")
	})

}
