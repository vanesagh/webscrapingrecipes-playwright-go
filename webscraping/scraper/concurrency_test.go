package scraper

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "waat:// fuhurterwe.geds" {
		return false
	}
	return true
}

func TestConcurrency(t *testing.T) {
	websites := []string{
		"http://google.com",
		"waat://furhurterwe.gde",
		"http://blog.gypsydave5.com",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"waat://furhurterwe.gde":     false,
		"http://blog.gypsydave5.com": true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)
	if reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}

}
