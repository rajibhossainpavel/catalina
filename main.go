package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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

	// Find the review items

	doc.Find("tbody").Each(func(i int, s *goquery.Selection) {
		writableString := "\n"
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			writableString = writableString + "," + s.Text()

		})
		fmt.Printf("Review %d: %s\n", i, writableString)
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
func main() {
	GetData("dse-raw-data.txt")
	GetNewFile("dse-raw-data.txt", "dse-raw-new-data.txt", "Helpdesk for NRB")
	GetNewFile("dse-raw-new-data.txt", "dse-raw-new-2-data-.txt", "1JANATAMF")
	WriteNewFile("dse-raw-new-2-data-.txt", "dse-raw-new-3-data-.txt", "If YCP is available")
	//fmt.Println("%b", success)

}
