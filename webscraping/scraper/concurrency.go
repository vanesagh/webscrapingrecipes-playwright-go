package scraper

import "fmt"

type WebsiteChecker func(string) bool

type Job struct {
	URL string
	err error
	//Success bool
}

type result struct {
	string
	bool
}

func ScrapeInParallel(urls interface{}) {
	urlsList := urls.([]interface{})
	jobs := make(chan Job, len(urlsList[8:10]))
	//results := make(chan Job)
	//ch := make(chan struct{})

	for _, url := range urlsList[8:10] {
		u := url.(string)
		go func(u string) {
			var job Job
			job.err = ScrapeForRecipe(u)
			job.URL = u
			jobs <- job
			fmt.Println(u)

		}(u)

	}
	//len(urlsList[:1])
	for i := 0; i < len(urlsList[8:10]); i++ {
		job := <-jobs
		if job.err != nil {
			fmt.Println(job.err)
		}
	}
	//close(jobs)
	//close(results)

}

func CheckWebsites(wc WebsiteChecker, urls interface{}) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	ffs := urls.([]interface{})

	for _, url := range ffs {
		s := url.(string)

		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(s)

	}
	for i := 0; i < len(ffs); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results

}
