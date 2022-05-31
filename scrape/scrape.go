package scrape

import (
	"bufio"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func GetDate() string {
	// Request the HTML page.
	res, err := http.Get("https://dsebd.org/latest_share_price_scroll_l.php")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	targetDate := ""

	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		targetDate = strings.ToLower(s.Text())
		if strings.Contains(targetDate, "latest share price") {
			targetDate = strings.ReplaceAll(targetDate, "latest share price", "")
			targetDate = strings.ReplaceAll(targetDate, "on", "")
			targetDate = strings.TrimSpace(targetDate)
			targetDate = strings.Title(targetDate)
			targetDate = strings.ReplaceAll(targetDate, "At", "at")
			targetDate = strings.ReplaceAll(targetDate, " Pm", "pm")
			explodedDate := strings.Split(targetDate, "at")
			targetDate = strings.TrimSpace(explodedDate[0])
			explodedDate = strings.Split(targetDate, ",")
			dateString := strings.TrimSpace(explodedDate[len(explodedDate)-1])
			dateString += "-"
			monthDate := strings.TrimSpace(explodedDate[0])
			explodedDate = strings.Split(monthDate, " ")
			month := strings.TrimSpace(explodedDate[0])
			switch month {
			case "Jan":
				month = "01"
				break
			case "Feb":
				month = "02"
				break
			case "Mar":
				month = "03"
				break
			case "Apr":
				month = "04"
				break
			case "May":
				month = "05"
				break
			case "Jun":
				month = "06"
				break
			case "Jul":
				month = "07"
				break
			case "Aug":
				month = "08"
				break
			case "Sep":
				month = "09"
				break
			case "Oct":
				month = "10"
				break
			case "Nov":
				month = "11"
				break
			case "Dec":
				month = "11"
				break

			}
			dateString += month
			dateString += "-"
			date := strings.TrimSpace(explodedDate[len(explodedDate)-1])
			dateString += date

			targetDate = dateString

		}
	})

	return targetDate

}

func GetData(path string) {
	handle, err := os.Create(path)
	check(err)
	writeBuffer := bufio.NewWriter(handle)
	defer handle.Close()
	// Request the HTML page.
	res, err := http.Get("https://dsebd.org/latest_share_price_scroll_l.php")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		writableString := "\n"
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			writableString = writableString + "," + s.Text()

		})
		//fmt.Printf("Review %d: %s\n", i, writableString)
		_, err := writeBuffer.WriteString(writableString)
		check(err)

	})
	writeBuffer.Flush()
}

func GetNewFile(sourcePath string, destinationPath string, searchString string) bool {

	sourceHandle, err := os.Open(sourcePath)
	if err != nil {
		//
	}

	destinationHandle, err := os.Create(destinationPath)
	check(err)
	writableBuffer := bufio.NewWriter(destinationHandle)

	defer sourceHandle.Close()
	defer destinationHandle.Close()
	scanner := bufio.NewScanner(sourceHandle)

	line := 1
	success := false

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchString) {
			success = true
			_, err := writableBuffer.WriteString(",1," + "\n")
			check(err)
		}
		if success {
			_, err := writableBuffer.WriteString(scanner.Text() + "\n")
			check(err)
		}
		line++
	}
	writableBuffer.Flush()

	if err := scanner.Err(); err != nil {
		// Handle the error
	}
	return success
}
func WriteNewFile(sourcePath string, destinationPath string, searchString string) bool {

	sourceHandle, err := os.Open(sourcePath)
	if err != nil {
		//
	}

	destinationHandle, err := os.Create(destinationPath)
	check(err)
	writableBuffer := bufio.NewWriter(destinationHandle)

	defer sourceHandle.Close()
	defer destinationHandle.Close()
	scanner := bufio.NewScanner(sourceHandle)

	line := 1

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchString) {
		} else {
			_, err := writableBuffer.WriteString(scanner.Text() + "\n")
			check(err)
		}
		line++
	}
	writableBuffer.Flush()

	if err := scanner.Err(); err != nil {
		// Handle the error
	}

	return true
}

func WriteCSVFile(sourcePath string, destinationPath string) bool {

	sourceHandle, err := os.Open(sourcePath)
	if err != nil {
		//
	}

	destinationHandle, err := os.Create(destinationPath)
	check(err)
	writableBuffer := bufio.NewWriter(destinationHandle)

	defer sourceHandle.Close()
	defer destinationHandle.Close()
	scanner := bufio.NewScanner(sourceHandle)

	line := 1
	name := ""
	for scanner.Scan() {

		text := strings.TrimSpace(scanner.Text())
		if strings.HasSuffix(text, ",") {
			name = ""
		} else if IsLetter(text) {
			name = text
		} else {
			_, err := writableBuffer.WriteString(name + text + "\n")
			check(err)
		}
		line++
	}
	writableBuffer.Flush()

	if err := scanner.Err(); err != nil {
		// Handle the error
	}

	return true
}
