package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type visit struct {
	url   string
	links []string
	depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	visits := make(map[string]visit)
	ch := make(chan visit)

	go crawl(url, depth, fetcher, ch)

	for started := 1; started > 0; {
		v := <-ch
		visits[v.url] = v
		started--

		if v.depth > 0 {
			for _, next := range v.links {
				if _, visited := visits[next]; !visited {
					go crawl(next, v.depth, fetcher, ch)
					started++
				}
			}
		}
	}
}

func crawl(url string, depth int, fetcher Fetcher, ch chan visit) {
	body, links, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		ch <- visit{url, []string{}, depth - 1}
	} else {
		fmt.Printf("found: %s %q\n", url, body)
		ch <- visit{url, links, depth - 1}
	}
}

func main() {
	Crawl("http://golang.org/", 40, fetcher)
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
