package main

import "fmt"

// Useful packages:
// - flags
// - csv
// - os (open, read file)
// - time, channels & goroutine for timer

// Technical requirements:
// - Read a CSV file ( default filename: problems.csv)
// - CSV file format: <question>, <answer>
// - CSV files may have questions with multiple commas like: "what 2+2, sir?",4
// - Quizzes will have less than 100 questions and will have single word/number answers
// - User should be able to pass any CSV filename via a flag

// Functional requirements:
// - Keep count of correct & incorrect answers
// - Show a quiz to a user
// - Regardless of the answer, display next question immediately after
// - At the end of the quiz, display the total number of correct answers and total questions

func main() {
	fmt.Printf("Yellow !")
}
