package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	question         string
	answerOptionA    string
	answerOptionB    string
	answerOptionC    string
	answerOptionD    string
	correctionOption string
}

func main() {
	questions, err := readQuestionsFromCSV("questions.csv")
	if err != nil {
		fmt.Println("Error reading questions from CSV file:", err)
		return
	}

	fmt.Println("Welcome to the quiz game!")
	score := 0
	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.question)
		fmt.Println("a. ", q.answerOptionA)
		fmt.Println("b. ", q.answerOptionB)
		fmt.Println("c. ", q.answerOptionC)
		fmt.Println("d. ", q.answerOptionD)

		fmt.Print("Please enter your option (a, b, c or d): ")
		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		if option == q.correctionOption {
			fmt.Println("Correct answer!")
			score++
		} else {
			fmt.Println("Incorrect answer.")
		}
	}

	fmt.Printf("\nYou scored %d out of %d.\n", score, len(questions))
}

func readQuestionsFromCSV(filename string) ([]quiz, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	questions := make([]quiz, len(records))
	for i, record := range records {
		questions[i] = quiz{
			question:         record[0],
			answerOptionA:    record[1],
			answerOptionB:    record[2],
			answerOptionC:    record[3],
			answerOptionD:    record[4],
			correctionOption: record[5],
		}
	}

	return questions, nil
}
