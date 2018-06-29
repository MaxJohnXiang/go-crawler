package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds{
		requests  = append(requests,r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests =  requests[1:]
		parseResult ,err := worker(r)
		if err != nil {
			continue
		}
		requests = append(
			requests,
			parseResult.Requests...
		)
		for _, item :=range parseResult.Items {
			log.Printf("Got Item %v",  item)
		}
	}
}

func worker(r Request) (ParseResult, error ){
		body ,err := fetcher.Fetch(r.Url)
		log.Printf("Fetching %s", r.Url)
		if err != nil {
			log.Printf("Fetcher error :" + "fetching url %s: %v" ,r.Url, err)
			return ParseResult{}, err
		}
		parseResult := r.ParserFunc(body)
	return parseResult, err
}
