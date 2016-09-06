package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
)

var path string

func init() {
	flag.StringVar(&path, "path", "./example.csv", "Path to the CSV file")
	flag.Parse()
}
func main() {
	if path == "" {
		log.Fatal("No path specified")
	}
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
