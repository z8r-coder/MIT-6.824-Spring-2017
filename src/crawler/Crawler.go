package crawler

import (
	"fmt"
	"sync"
)

func CrawSerial(url string, fetcher Fetcher, fetched map[string]bool)  {
	if fetched[url] {
		return
	}
	fetched[url] = true
	body, urls, err := fetcher.Fetche(url)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _,u := range urls {
		CrawSerial(u, fetcher, fetched)
	}

	return 
}

type fetchState struct {
	mu sync.Mutex
	fetched map[string]bool
}

func (f *fetchState) CheckAndMark(url string) bool {
	if f.fetched[url] {
		return true
	}
	f.fetched[url] = true
	return false
}

func mkFetchState() *fetchState {
	f := &fetchState{}
	f.fetched = make(map[string]bool)
	return f
}

func CrawlConcurrentMutex(url string, fetcher Fetcher, f *fetchState)  {
	if f.CheckAndMark(url) {
		return
	}
	body, urls, err := fetcher.Fetche(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("found: %s %q\n", url, body)
	var done sync.WaitGroup
	for _, u := range urls{
		done.Add(1)
		go func(u string) {
			defer done.Done()
			CrawlConcurrentMutex(u, fetcher, f)
		}(u)
	}
	done.Wait()
	return
}

func dofetch(url1 string, ch chan []string, fetcher Fetcher)  {
	body, urls, err := fetcher.Fetche(url1)
	if err != nil {
		fmt.Println(err)
		ch <- []string{}
	} else {
		fmt.Printf("found: %s %q\n", url1, body)
		ch <- urls
	}
}

func master(ch chan []string, fetcher Fetcher)  {
	n := 1
	fetched := make(map[string]bool)
	for urls := range ch {
		for _, u:= range urls{
			if _, ok := fetched[u]; ok == false {
				fetched[u] = true
				n += 1
				go dofetch(u, ch, fetcher)
			}
		}
		n -= 1
		if n == 0 {
			break
		}
	}
}

func CrawlConcurrentChannel(url string, fetcher Fetcher)  {
	ch := make(chan []string)
	go func() {
		ch <- []string{url}
	}()
	master(ch, fetcher)
}

func main()  {
	fmt.Println("==========Serial==========\n")
	CrawSerial("http://golang.org/", fetcher, make(map[string]bool))

	fmt.Printf("====== ConcurrentMutex =====\n")
	CrawlConcurrentMutex("http://golang.org/", fetcher, mkFetchState())

	fmt.Println("===== ConcurrentChannel ======\n")
	CrawlConcurrentChannel("http://golang.org/", fetcher)
}

type Fetcher interface {
	Fetche(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string] *fakeResult

type fakeResult struct {
	body string
	urls[] string
}

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