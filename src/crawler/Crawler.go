package crawler

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetche(url string) (body string, urls []string, err error)
}

type fakeFetcher map[string] *fakeResult

type fakeResult struct {
	body string
	urls[] string
}


func main() {
	fmt.Println(111)
}
