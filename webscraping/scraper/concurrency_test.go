package scraper

import (
	"fmt"
	"strings"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	return strings.HasPrefix(url, "https://theviewfromgreatisland.com/")
}

func TestConcurrency(t *testing.T) {

	t.Run("check json file list", func(t *testing.T) {
		urlsList, _ := OpenFile("../muffins.json")
		List := urlsList.([]interface{})
		fmt.Printf("%T", List)
		got := CheckWebsites(mockWebsiteChecker, List)
		fmt.Println(got)

	})

	t.Run("check function", func(t *testing.T) {
		urlsList, _ := OpenFile("../muffins.json")
		List := urlsList.([]interface{})
		ScrapeInParallel(List)

	})

}
