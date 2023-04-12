package scraper

import (
	"testing"
)

func assertError(t testing.TB, err error) {
	t.Helper()
	if err == nil {

		t.Error("wanted error but didn't get one")
	}
}

func TestFileHandling(t *testing.T) {

	t.Run("wrong extension", func(t *testing.T) {
		path := "recipes.doc"
		_, err := OpenFile(path)
		assertError(t, err)

	})

	t.Run("file not found", func(t *testing.T) {
		path := "recipes.json"
		_, err := OpenFile(path)
		assertError(t, err)
	})

	/*t.Run("open file", func(t *testing.T) {
		listUrls, _ := OpenFile("../muffins.json")
		listOfUrls := listUrls.([]interface{})
		fmt.Printf("%T", listOfUrls)
		for i, r := range listOfUrls {
			//rs := r.(string)
			fmt.Printf("%d, %T", i, r)

		}

	})*/

}
