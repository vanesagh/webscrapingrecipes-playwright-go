package scraper

import (
	"log"

	"github.com/playwright-community/playwright-go"
)

func Initialize() (playwright.Page, *playwright.Playwright, playwright.Browser) {
	pw, err := playwright.Run()
	AssertErrorToNil("could not start playwright: ", err)

	browser, err := pw.Chromium.Launch()
	AssertErrorToNil("could not launch browser: ", err)

	page, err := browser.NewPage()
	page.SetDefaultTimeout(10000) //needed more time. Otherwise got nothing
	AssertErrorToNil("could not create page: ", err)

	return page, pw, browser

}

func Close(pw *playwright.Playwright, browser playwright.Browser) {
	if err := browser.Close(); err != nil {
		log.Fatalf("could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("could not stop Playwright: %v", err)
	}

}
