package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type urlsMap map[string][]string

func (u urlsMap) saveToMap(set urlRecipeSet, recipe string) {
	var listURLs []string

	for k := range set {
		listURLs = append(listURLs, k)
	}
	u[recipe] = listURLs
}

func (u urlsMap) convertToJSON() []byte {
	jsonFormat, err := json.Marshal(u)
	AssertErrorToNil("error converting to json", err)
	fmt.Println(string(jsonFormat))
	return jsonFormat
}

func saveJsonToFile(jsonFormat []byte, recipeName string) {
	fileName := recipeName + ".json"
	err := ioutil.WriteFile(fileName, jsonFormat, 0644)
	AssertErrorToNil("error saving file", err)

}
