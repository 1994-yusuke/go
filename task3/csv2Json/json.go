package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.String("", "", "default")
	flag.Parse()
	files := flag.Args()
	values := files[0]
	if err := readJson(values); err != nil {
		fmt.Printf("error: %s\n", err)
	}
}

func readJson(value string) error {

	type Items struct {
		Year    int
		Name    string
		Company string
	}

	csvItem, err := os.Open(value)

	if err != nil {
		return err
	}
	defer csvItem.Close()

	reader := csv.NewReader(csvItem)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var oneRecord Items
	var allRecords []Items

	for _, each := range csvData {
		oneRecord.Year, _ = strconv.Atoi(each[0])
		oneRecord.Name = each[1]    // need to cast integer to string
		oneRecord.Company = each[2] // need to cast integer to string
		// need to cast integer to string

		allRecords = append(allRecords, oneRecord)
	}

	jsons, err := json.Marshal(allRecords) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(jsons))
	return err
}
