package main

import (
	"Crawler2db/Crawler"
	"fmt"
)

func main() {
	MetroInfo := Crawler.Webcrawler()
	fmt.Print(MetroInfo)
}
