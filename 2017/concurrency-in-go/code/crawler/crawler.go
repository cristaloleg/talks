package code

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
func (b *BigBrother) NoNoJustVisiting(url string, f func(string)) {
	links := b.Extract(b.Process(url))
	size := len(links)
	choker := make(chan struct{}, 32)
	wg := sync.WaitGroup{}

	for i := 0; i < size; i++ {
		choker <- struct{}
		wg.Add(1)

		go func(url string, wg sync.WaitGroup, choker <-chan struct{}) {
			f(url)
			
			links := b.Extract(b.Process(url))
			// whatever.DoWith(links)

			<-choker
			wg.Done()
		}(links[i], &wg, choker)
	}
	wg.Wait()
}

// END_CRAWL OMIT
