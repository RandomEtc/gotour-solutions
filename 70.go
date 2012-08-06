package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (result *fakeResult, err error)
}

var fetched = map[string]bool{}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, resultch chan *fakeResult, errch chan error) {
	if depth <= 0 {
		return
	}
	res, err := fetcher.Fetch(url)
	if err != nil {
		errch <- err
		return
	}
	resultch <- res
	for _, u := range res.urls {
		if !fetched[u] {
			fetched[u] = true
			go Crawl(u, depth-1, fetcher, resultch, errch)
		}
	}
	return
}

func main() {
	resultch := make(chan *fakeResult)
	errch := make(chan error)
	fetched["http://golang.org/"] = true
	go Crawl("http://golang.org/", 4, fetcher, resultch, errch)
	for {
		select {
		case result := <- resultch:
			//fmt.Printf("found: %s %q\n", url, res.body)
			fmt.Printf("found: %q\n", result.body)
		case err := <- errch:
			fmt.Println(err)
		}
	}
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls     []string
}

func (f *fakeFetcher) Fetch(url string) (*fakeResult, error) {
	if res, ok := (*f)[url]; ok {
		return res, nil
	}
	return nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
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