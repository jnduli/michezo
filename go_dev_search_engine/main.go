package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

type WebsiteRss struct {
	name    string
	website string
	rssFeed string
}

func readCsvFile(filePath string) []WebsiteRss {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var websiteRecords []WebsiteRss
	for idx, record := range records {
		if idx == 0 {
			continue
		}
		local := WebsiteRss{
			name:    record[0],
			website: record[1],
			rssFeed: record[2],
		}
		websiteRecords = append(websiteRecords, local)
	}
	return websiteRecords
}

func readFeed(feedUrl string) {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedUrl)
	fmt.Println(feed.Title)
}

func main() {
	records := readCsvFile("eng_websites.csv")
	fmt.Println(records)
	fmt.Println(records[0].rssFeed)
	readFeed(records[0].rssFeed)
}
