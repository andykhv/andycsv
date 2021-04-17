package main

import (
	"fmt"
	"math"
	"strconv"
)

type ChaseCheckingStatements [][]string
type ChaseCreditStatements [][]string
type CitiCreditStatements [][]string

/*
Converts Chase Checking Statements to Andy's CSV Format
Chase Checking Statement Format:
Details,Posting Date,Description,Amount,Type,Balance,Check or Slip #
*/
func (statements ChaseCheckingStatements) Convert() *[][]string {
	const (
		dateIndex        = 1
		descriptionIndex = 2
		amountIndex      = 3
	)

	return convert(statements, dateIndex, descriptionIndex, amountIndex)
}

/*
Converts Chase Credit Statements to Andy's CSV Format
Chase Credit Statement Format:
Transaction Date,Post Date,Description,Category,Type,Amount,Memo
*/
func (statements ChaseCreditStatements) Convert() *[][]string {
	const (
		dateIndex        = 0
		descriptionIndex = 2
		amountIndex      = 5
	)

	return convert(statements, dateIndex, descriptionIndex, amountIndex)
}

/*
Converts Citi Credit Statements to Andy's CSV Format
Citi Credit Statement Format:
Status,Date,Description,Debit,Credit,Member Name
*/
func (statements CitiCreditStatements) Convert() *[][]string {
	const (
		dateIndex        = 1
		descriptionIndex = 2
		amountIndex      = 3
	)

	return convert(statements, dateIndex, descriptionIndex, amountIndex)
}

func convert(statements [][]string, dateIndex int, descriptionIndex int, amountIndex int) *[][]string {
	andyStatements := make([][]string, len(statements))
	andyStatement := make([]string, len(statements)*len(header))

	for i := range andyStatements {
		date := statements[i][dateIndex]
		description := statements[i][descriptionIndex]
		amount := convertAmountToAbsolute(statements[i][amountIndex])
		category, err := GetCategory(&description, &amount)

		if err == nil {
			andyStatement[0] = date
			andyStatement[1] = description
			andyStatement[2] = category
			andyStatement[3] = amount

			andyStatements[i], andyStatement = andyStatement[:4], andyStatement[4:]
		}
	}

	return &andyStatements
}

func convertAmountToAbsolute(amount string) string {
	amountFloat, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		fmt.Printf("amount %s cannot be converted to float\n", amount)
	}

	amountFloat = math.Abs(amountFloat)
	return strconv.FormatFloat(amountFloat, 'f', -1, 64)
}
