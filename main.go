package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		log.Fatal(err)
	}

	csvFile := csv.NewReader(file)

	data, err := csvFile.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range data {
		fmt.Println(d)
	}
}
