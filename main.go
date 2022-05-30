package main

import (
	"catalina/scrape"
	"fmt"
	"log"
	"os"
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

func main() {
	dirERxists, _ := exists("data")
	if !dirERxists {
		if err := os.Mkdir("data", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("Processing Data.\n")
	scrape.GetData("data/data-1.txt")
	scrape.GetNewFile("data/data-1.txt", "data/data-2.txt", "Helpdesk for NRB")
	scrape.GetNewFile("data/data-2.txt", "data/data-3.txt", "1JANATAMF")
	scrape.WriteNewFile("data/data-3.txt", "data/data-4.txt", "If YCP is available")
	scrape.WriteCSVFile("data/data-4.txt", "data.csv")
	returnedDate := scrape.GetDate()
	fmt.Println("returnedDate: ", returnedDate)
	e := os.RemoveAll("data")
	if e != nil {
		log.Fatal(e)
	}
}
