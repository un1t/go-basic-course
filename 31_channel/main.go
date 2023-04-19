package main

import (
	"fmt"
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

	ch := make(chan string)

	go func() {
		for _, url := range urls {
			downloadFile(url)
			ch <- "ok"
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}

}
