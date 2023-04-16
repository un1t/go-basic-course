package main

import (
	"fmt"
	"sync"
	"time"
)

func downloadFile(url string) {
	fmt.Println("starting:" + url)
	time.Sleep(time.Second * 1)
	fmt.Println("done:" + url)
}

func main() {
	urls := []string{
		"http://example.com/file1.mp4",
		"http://example.com/file2.mp4",
		"http://example.com/file3.mp3",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			downloadFile(url)
		}(url)
	}

	wg.Wait()
}
