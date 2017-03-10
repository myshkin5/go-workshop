package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type crawler struct {
	urlsVisited map[string]bool
	lock        sync.Mutex
	wg          sync.WaitGroup
}

func New() *crawler {
	return &crawler{urlsVisited: make(map[string]bool)}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func (c *crawler) Crawl(url string, depth int, fetcher Fetcher) {
	c.wg.Add(1)
	c.urlsVisited[url] = true
	c.crawl(url, depth, fetcher)
}

func (c *crawler) crawl(url string, depth int, fetcher Fetcher) {
	defer c.wg.Done()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q with %d urls\n", url, body, len(urls))
	c.lock.Lock()
	defer c.lock.Unlock()
	for _, u := range urls {
		if !c.urlsVisited[u] {
			c.urlsVisited[u] = true
			c.wg.Add(1)
			go c.crawl(u, depth-1, fetcher)
		}
	}
}

func (c *crawler) Wait() {
	c.wg.Wait()
}

func main() {
	crawler := New()
	crawler.Crawl("http://golang.org/", 4, fetcher)
	crawler.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
