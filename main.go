package main

import (
	"catalina/scrape"
	"fmt"
	"os"
	"strings"
	"time"
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

func ParseData() (bool, error) {
	today, err := GetToday()
	if err != nil {
		return false, err
	}
	runCondition := true
	tradeDate := ""
	for {
		tradeDate, err := scrape.GetDate()
		if err == nil {
			if tradeDate != "" {
				runCondition = false
			}
		}
		if !runCondition {
			break
		}
	}
	if today != tradeDate {
		fmt.Printf("Processing Data.\n")
		dataRunCondition := true
		dataSucess := false
		for {
			dataSucess, err := scrape.GetData("data/data-1.txt")
			if err == nil {
				if dataSucess == true {
					dataRunCondition = false
				}
			}
			if !dataRunCondition {
				break
			}
		}
		if dataSucess {
			newFile, err := scrape.GetNewFile("data/data-1.txt", "data/data-2.txt", "Helpdesk for NRB")
			if err == nil {
				if newFile {

					newFile2, err := scrape.GetNewFile("data/data-2.txt", "data/data-3.txt", "1JANATAMF")
					if err == nil {
						if newFile2 {
							newFile3, err := scrape.WriteNewFile("data/data-3.txt", "data/data-4.txt", "If YCP is available")
							if err == nil {
								if newFile3 {
									newFile4, err := scrape.WriteCSVFile("data/data-4.txt", "data.csv")
									if err == nil {
										if newFile4 {
											return true, nil

										}
									}

								}
							}

						}
					}

				}
			}

		}
	}
	return true, nil
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

func main() {
	dirSuccess, err := CreateDir("data")
	if err == nil {
		if dirSuccess {
			success, err := ParseData()
			if err == nil {
				if success {
					/*e := os.RemoveAll("data")
					if e != nil {
						log.Fatal(e)
					}*/
				}

			}
		}
	}
}
