package probe

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Probe struct {
	workers int
	url     string
}

func NewProbe(worker int, url string) *Probe {
	return &Probe{
		workers: worker,
		url:     url,
	}
}

func (p *Probe) GetFileSize() (int, error) {
	var size = -1

	client := &http.Client{}
	req, err := http.NewRequest("HEAD", p.url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if header, ok := resp.Header["Content-Length"]; ok {
		fileSize, err := strconv.Atoi(header[0])
		if err != nil {
			log.Fatal("File size could not be determined : ", err)
		}

		size = fileSize / p.workers
	} else {
		log.Fatal("File size was not provided!")
		return size, fmt.Errorf("file size was not provided.")
	}
	return size, nil
}
