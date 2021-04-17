package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFile, csvType, outputFile, err := getFlags()
	if err != nil {
		fmt.Println(err)
		return
	}

	statementFile, err := os.Open(csvFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer statementFile.Close()

	csvReader := csv.NewReader(statementFile)
	records, _ := csvReader.ReadAll()
	var andyRecords *[][]string
	switch {
	case csvType == "chase-checking":
		andyRecords = ChaseCheckingStatements(records[1:]).Convert()
	case csvType == "chase-credit":
		andyRecords = ChaseCreditStatements(records[1:]).Convert()
	case csvType == "citi-credit":
		andyRecords = CitiCreditStatements(records[1:]).Convert()
	}

	andyFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer andyFile.Close()

	csvWriter := csv.NewWriter(andyFile)
	csvWriter.Write(header)
	csvWriter.WriteAll(*andyRecords)
	csvWriter.Flush()
}

func getFlags() (csvFile string, csvType string, outputFile string, err error) {
	flag.StringVar(&csvFile, "csv", "", "must specify file to parse")
	flag.StringVar(&csvType, "type", "", "must specify initial csv format, options: 'chase-checking', 'chase-credit', 'citi-credit'")
	flag.StringVar(&outputFile, "output", "output.csv", "specify file for output csv")

	flag.Parse()

	if csvFile == "" {
		return "", "", "", errors.New("must specify file to parse")
	}

	if !isCsvTypeValid(&csvType) {
		return "", "", "", errors.New("must specify valid csv type, options: 'chase-checking', 'chase-credit', 'citi-credit'")
	}

	return
}

func isCsvTypeValid(csvType *string) bool {
	return *csvType == "chase-checking" || *csvType == "chase-credit" || *csvType == "citi-credit"
}
