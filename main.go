package main

import (
	"Crawler2db/crawler"
	"Crawler2db/postgres"
	"fmt"
)

func main() {
	MetroInfo := crawler.Webcrawler()
	// fmt.Print(MetroInfo)
	err := postgres.Update(MetroInfo)
	fmt.Print(err)
}
