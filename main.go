package main

import (
	"fmt"
	"golang-test-task/task"
)

func main() {

	fmt.Println("======================================")
	fmt.Println("         GoLang Test Task")
	fmt.Println("======================================")
	generatedString := task.Generate(true)
	fmt.Println("The generated string: ", generatedString)
	fmt.Println("======================================")
	fmt.Println("Test Validity: ", task.TestValidity(generatedString))
	fmt.Println("======================================")
	fmt.Println("Average Number: ", task.AverageNumber(generatedString))
	fmt.Println("======================================")
	fmt.Println("Whole Story: ", task.WholeStory(generatedString))
	fmt.Println("======================================")
	fmt.Print("Story Stats: ")
	fmt.Println(task.StoryStats(generatedString))
	fmt.Println("======================================")
}
