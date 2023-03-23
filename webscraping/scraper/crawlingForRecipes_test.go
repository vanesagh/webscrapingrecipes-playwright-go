package scraper

import (
	"testing"
)

func TestCrawlingForRecipes(t *testing.T) {
	t.Run("not found", func(t *testing.T) {
		recipe := "yudane"
		CrawlForWebPageRecipes(recipe)

	})

}
