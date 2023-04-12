package scraper

import (
	"testing"
)

func TestScrapeRecipe(t *testing.T) {
	t.Run("handle wrong url", func(t *testing.T) {
		urlRecipe := "https://www.google.com"
		err := ScrapeForRecipe(urlRecipe)

		assertError(t, err)

	})

	t.Run("handle recipes with one process", func(t *testing.T) {
		err := ScrapeForRecipe("https://theviewfromgreatisland.com/perfect-blueberry-muffins/#cls-video-container-zLMOOuwt")
		assertError(t, err)
	})

}
