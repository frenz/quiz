package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type problemData struct {
	Question string
	Answer   string
}

func printError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func readCsvLines(filename string) [][]string {
	csvFile, err := os.Open(filename)
	printError(err)
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	printError(err)
	return csvLines
}
func parseLinesToProblems(csvLines [][]string) []problemData {
	var problems []problemData
	for _, line := range csvLines {
		problem := problemData{
			Question: line[0],
			Answer:   line[1],
		}
		problems = append(problems, problem)
	}
	return problems
}

func askUser(problems []problemData) int {
	correct := 0
	for _, problem := range problems {
		answer := ""
		fmt.Printf("%s ?\n", problem.Question)
		fmt.Scanf("%s", &answer)
		if answer == problem.Answer {
			correct++
		}
	}
	return correct
}
func main() {
	filename := "problems.csv"
	csvLines := readCsvLines(filename)
	problems := parseLinesToProblems(csvLines)
	correct := askUser(problems)
	fmt.Printf("User answered %d correct %d not correct", correct, len(problems)-correct)
}
