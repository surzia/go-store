package download

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Downloader struct {
	result  chan Part
	size    int
	workers int
}

func NewDownloader(result chan Part, size, worker int) *Downloader {
	d := &Downloader{
		result:  result,
		size:    size,
		workers: worker,
	}

	return d
}

func (d *Downloader) Download(index int, url string) {
	client := &http.Client{}

	// calculate offset by multiplying
	// index with size
	start := index * d.size

	// Write data range in correct format
	// I'm reducing one from the end size to account for
	// the next chunk starting there
	dataRange := fmt.Sprintf("bytes=%d-%d", start, start+d.size-1)

	// if this is downloading the last chunk
	// rewrite the header. It's an easy way to specify
	// getting the rest of the file
	if index == d.workers-1 {
		dataRange = fmt.Sprintf("bytes=%d-", start)
	}
	log.Println(dataRange)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// TODO: restart download
		return
	}
	req.Header.Add("Range", dataRange)
	resp, err := client.Do(req)
	if err != nil {
		// TODO: restart download
		return
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO: restart download
		return
	}

	d.result <- Part{Index: index, Data: body}
}

func (d *Downloader) Merge(filename string) error {
	log.Println("start to merge data")

	parts := make([][]byte, d.workers)
	counter := 0
	for part := range d.result {
		counter++
		parts[part.Index] = part.Data
		if counter == d.workers {
			break
		}
	}
	log.Println("sort data as original order")

	file := []byte{}
	for _, part := range parts {
		file = append(file, part...)
	}
	log.Println("write data into buffer array")
	err := ioutil.WriteFile(filename, file, 0777)
	return err
}
