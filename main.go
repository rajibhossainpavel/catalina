package main

import (
	"catalina/mongodb"
	"catalina/scrape"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetToday() (string, error) {
	t := time.Now()
	location, err := time.LoadLocation("Asia/Dhaka")
	if err != nil {
		return "", nil
	}
	now := t.In(location).Format(time.RFC3339)
	explodedDate := strings.Split(now, "T")
	dateString := strings.TrimSpace(explodedDate[0])
	return dateString, nil
}

func CreateDir(path string) (bool, error) {
	dirERxists, _ := exists(path)
	if !dirERxists {
		if err := os.Mkdir(path, os.ModePerm); err != nil {
			return false, err
		}
	}
	return true, nil
}

func ParseData(url string) (bool, error) {
	var doc *goquery.Document
	doc = nil
	runCondition := true
	for {
		doc, _ = scrape.GetDocument(url)
		if doc != nil {
			runCondition = false
		}
		if !runCondition {
			break
		}
	}

	if doc != nil {
		tradeDate, _ := scrape.GetDate(doc)
		today, _ := GetToday()
		if today != "" && tradeDate == today {
			result, _ := scrape.GetData(doc, "data/data-1.txt")
			if result {
				if result {
					result, _ = scrape.GetNewFile("data/data-1.txt", "data/data-2.txt", "Helpdesk for NRB")
					if result {
						result, _ = scrape.GetNewFile("data/data-2.txt", "data/data-3.txt", "1JANATAMF")
						if result {
							result, _ = scrape.WriteNewFile("data/data-3.txt", "data/data-4.txt", "If YCP is available")
							if result {
								result, _ = scrape.WriteCSVFile("data/data-4.txt", "data.csv")
								if result {
									return true, nil

								}
							}
						}
					}
				}
			}
		}
	}
	return false, nil
}

func main() {
	/*dirSuccess, _ := CreateDir("data")
	if dirSuccess {
		result, _ := ParseData("https://dsebd.org/latest_share_price_scroll_l.php")
		if result {
			e := os.RemoveAll("data")
			if e != nil {
				log.Fatal(e)
			} else {


			}
		}
	}*/
	mongodb.Connect()
}
