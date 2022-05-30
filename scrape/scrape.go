package scrape

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

func GetDate() {
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

	doc.Find("h2").Each(func(i int, s *goquery.Selection) {
		targetDate := strings.ToLower(s.Text())
		if strings.Contains(targetDate, "latest share price") {
			targetDate = strings.ReplaceAll(targetDate, "latest share price", "")
			targetDate = strings.ReplaceAll(targetDate, "on", "")
			//
			//
			targetDate = strings.TrimSpace(targetDate)
			targetDate = strings.Title(targetDate)
			targetDate = strings.ReplaceAll(targetDate, "At", "at")
			targetDate = strings.ReplaceAll(targetDate, "Pm", "pm")
			fmt.Println(targetDate)
			//explodedDate := strings.Split(targetDate, ",")

			//explodedDate = explodedDate[:len(explodedDate)-1]
			//monthDate := strings.Join(explodedDate, "")
			//monthDate = monthDate + ", 2022"
			//fmt.Println(monthDate)

			date, error := time.Parse(targetDate, targetDate)

			if error != nil {
				fmt.Println(error)
				return
			}
			fmt.Println(date)
		}
	})
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
	fmt.Printf("Processing Data:")
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
func main() {
	if err := os.Mkdir("data", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	GetData("data/data-1.txt")
	GetNewFile("data/data-1.txt", "data/data-2.txt", "Helpdesk for NRB")
	GetNewFile("data/data-2.txt", "data/data-3.txt", "1JANATAMF")
	WriteNewFile("data/data-3.txt", "data/data-4.txt", "If YCP is available")
	WriteCSVFile("data/data-4.txt", "data.csv")

	e := os.RemoveAll("data")
	if e != nil {
		log.Fatal(e)
	}
}
