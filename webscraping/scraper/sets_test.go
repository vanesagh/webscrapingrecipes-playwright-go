package scraper

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	newSet := urlRecipeSet{}
	newSet.AddElement("https://")
	newSet.AddElement("https://")
	newSet.AddElement("https://")
	newSet.AddElement("http")
	for key := range newSet {
		fmt.Println(key)
	}

}
