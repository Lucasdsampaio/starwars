package api

import (
	"fmt"
	"net/http"
)

func AsyncGet() ([]*HttpResponse, error) {
	ch := make(chan *HttpResponse)
	responses := []*HttpResponse{}
	pages := 6

	for pageID := 1; pageID <= pages; pageID++ {
		go func(pageID int) {
			url := fmt.Sprintf("https://swapi.dev/api/planets/?page=%d", pageID)
			response, err := http.Get(url)
			if err != nil {
				panic("Bad Gateway")
			}
			ch <- &HttpResponse{pageID, response, err}
		}(pageID)

	}

	for {
		select {
		case r := <-ch:
			responses = append(responses, r)
			if len(responses) == pages {
				return responses, nil
			}
		}
	}
}
