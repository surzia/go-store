package main

import (
	"flag"
	"log"
	"time"

	"go-store/applications/downloader/download"
	"go-store/applications/downloader/probe"
)

var (
	// to test internet
	url = flag.String("url", "http://212.183.159.230/512MB.zip", "download url")
	// number of goroutines to spawn for download.
	workers = flag.Int("worker", 5, "concurrent downloader number")
	// filename for downloaded file
	filename = flag.String("file", "data.zip", "downloaded filename")
)

func main() {
	flag.Parse()
	start := time.Now()
	probe := probe.NewProbe(*workers, *url)
	size, err := probe.GetFileSize()
	if err != nil {
		panic(err)
	}

	results := make(chan download.Part, *workers)
	downloader := download.NewDownloader(results, size, *workers)
	for i := 0; i < *workers; i++ {
		go downloader.Download(i, *url)
	}

	err = downloader.Merge(*filename)
	end := time.Now()
	if err != nil {
		panic(err)
	}

	log.Println("cost time: ", end.Sub(start))
}
