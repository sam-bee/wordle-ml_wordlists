package wordlemlwordlists

import _ "embed"

//go:embed data/wordlist-valid-guesses.csv
var validGuessesCSV string

//go:embed data/wordlist-valid-solutions.csv
var validSolutionsCSV string

func ValidGuessesCSV() string {
	return validGuessesCSV
}

func ValidSolutionsCSV() string {
	return validSolutionsCSV
}
