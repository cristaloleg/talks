package main

import "sync"

// START_DEFS OMIT
type Crawler interface {
	Extract(string) []string
	IsFinished(string) bool
	Process(string) string
}

type BigBrother struct {
	cache map[string]struct{}
}

// END_DEFS OMIT

// START_EXTRACT OMIT
func (b *BigBrother) Extract(url string) []string {
	// extract all links from url
	return nil
}

// END_EXTRACT OMIT

// START_EXISTS OMIT
func (b *BigBrother) IsFinished(url string) bool {
	_, ok := b.cache[url]
	return ok
}

// END_EXISTS OMIT

// START_PROCESS OMIT
func (b *BigBrother) Process(url string) string {
	// download content from url
	return ""
}

// END_PROCESS OMIT

// START_CRAWL OMIT
func (b *BigBrother) Crawl(url string, f func(string)) {
	links := b.Extract(b.Process(url))
	size := len(links)
	wg := sync.WaitGroup{}

	for i := 0; i < size; i++ {
		wg.Add(1)

		go func(url string) {
			f(url)
			
			wg.Done()
		}(links[i])
	}
	wg.Wait()
}
// END_CRAWL OMIT

// START_CRAWL2 OMIT
func (b *BigBrother) CrawlBounded(url string, f func(string)) {
	links := b.Extract(b.Process(url))
	size := len(links)
	wg := sync.WaitGroup{}
	choker := make(chan struct{}, 32)

	for i := 0; i < size; i++ {
		choker <- struct{}{}
		wg.Add(1)

		go func(url string, choker <-chan struct{}) {
			f(url)

			<-choker
			wg.Done()
		}(links[i], choker)
	}
	wg.Wait()
}

// END_CRAWL2 OMIT

func main() {
	
}
