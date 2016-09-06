package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: catcsv <pathtocsv>")
	}
	path := os.Args[1]
	fileReader, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not open file:", err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	r := csv.NewReader(fileReader)
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if i == 0 {
			table.SetHeader(record)
		} else {
			table.Append(record)
		}

		i++
	}
	table.Render()

}
