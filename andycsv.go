package main

import (
	"errors"
	"fmt"
)

var (
	header     = []string{"DATE", "DESCRIPTION", "CATEGORY", "AMOUNT"}
	categories = map[int]string{
		1:  "INCOME",
		2:  "BILLS",
		3:  "IRA",
		4:  "INVESTING",
		5:  "SAVINGS",
		6:  "NEEDS",
		7:  "WANTS",
		8:  "RESTAURANT",
		9:  "GROCERIES",
		10: "TRANSPORTATION",
		11: "FRIEND",
	}
)

func GetCategory(description *string, amount *string) (string, error) {
	printPrompt(description, amount)
	input := getInput()
	category := categories[input]

	if category == "" {
		return "", errors.New("category not specified")
	}

	return category, nil
}

func printPrompt(description *string, amount *string) {
	fmt.Printf("Statement description: %s\nStatement amount: %s\n", *description, *amount)
	printCategories()
	fmt.Printf("Choose category: ")
}

func printCategories() {
	fmt.Printf("Categories: ")
	for num, category := range categories {
		fmt.Printf("%s[%d], ", category, num)
	}
	fmt.Printf("\n")
}

func getInput() int {
	var input int
	fmt.Scanf("%d", &input)
	return input
}
