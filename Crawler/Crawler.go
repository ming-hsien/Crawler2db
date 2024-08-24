package crawler

import (
	// "fmt"
	"log"
	"net/http"

	// "sort"

	"github.com/PuerkitoBio/goquery"
)

type StationInfo struct {
	arrivalTime []string
}

type MetroInfo struct {
	stationName []string
	stationInfo map[string]StationInfo
	trainsCount int
}

func (MetroInfo) TableName() string {
	return "metroInfo"
}

func Webcrawler() map[string]MetroInfo {
	url := "https://ericyu.org/TaipeiMetroTime/lines/G-b-1,2,3,4,5.html"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("HTTP request error, Status Code：%d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	lines := map[string]MetroInfo{}
	currentLine := "G"
	stationName := []string{}
	trainsCount := 0

	doc.Find("table#stations").Each(func(index int, tableHtml *goquery.Selection) {
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			rowHtml.Find("td").Each(func(indexTd int, cellHtml *goquery.Selection) {
				cellText := cellHtml.Text()
				stationName = append(stationName, cellText)
			})
		})
	})

	stationInfo := make(map[string]StationInfo)
	// sort.Strings(lines[currentLine].stationName)
	// fmt.Print(lines[currentLine].stationName[0], "\n")
	count := 0
	doc.Find("table#timetable").Each(func(index int, tableHtml *goquery.Selection) {
		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {
			trainsCount = 0
			buf := stationInfo[stationName[count]]
			rowHtml.Find("td").Each(func(indexTd int, cellHtml *goquery.Selection) {
				cellText := cellHtml.Text()
				if cellText == "==" {
					buf.arrivalTime = append(buf.arrivalTime, "")
				} else {
					buf.arrivalTime = append(buf.arrivalTime, cellText)
				}
				trainsCount++
			})
			stationInfo[stationName[count]] = buf
			count++
		})
	})

	lines[currentLine] = MetroInfo{
		stationName: stationName,
		stationInfo: stationInfo,
		trainsCount: trainsCount,
	}

	return lines
}
