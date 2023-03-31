package scraper

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func OpenFile(filePath string) (interface{}, error) {

	ext := filepath.Ext(filePath)
	if ext != ".json" {
		return nil, errors.New("format file not supported")

	}
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(filePath)
	fileName = strings.TrimSuffix(fileName, ".json")
	//fmt.Println(fileName)
	var urlsList map[string]interface{}

	err = json.Unmarshal(fileContent, &urlsList)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("%T", urlsList[fileName])
	return urlsList[fileName], nil

}
