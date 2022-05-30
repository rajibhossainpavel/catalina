package main

import (
	"log"
	"os"
	"catalina/scrape"
)

func main() {
	if err := os.Mkdir("data", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	scrape.GetData("data/data-1.txt")
	scrape.GetNewFile("data/data-1.txt", "data/data-2.txt", "Helpdesk for NRB")
	scrape.GetNewFile("data/data-2.txt", "data/data-3.txt", "1JANATAMF")
	scrape.WriteNewFile("data/data-3.txt", "data/data-4.txt", "If YCP is available")
	scrape.WriteCSVFile("data/data-4.txt", "data.csv")

	e := os.RemoveAll("data")
	if e != nil {
		log.Fatal(e)
	}
}
