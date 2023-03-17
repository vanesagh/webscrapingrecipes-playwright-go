package scraper

type urlRecipeSet map[string]struct{}

var member struct{}

func (u urlRecipeSet) AddElement(url string) {
	u[url] = member

}
