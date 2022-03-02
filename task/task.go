package task

import (
	"strconv"
	"strings"
)

func TestValidity(str string) bool {

	splitArray := strings.Split(str, "-")
	for i := 0; i<len(splitArray); i++ {
		if modChecker(i) { 
			_, err := strconv.Atoi(splitArray[i])
			if err != nil {
				return false
			}
		} else {
			if len(splitArray[i]) == 0 {
				return false
			}
		}
	}
	return true
}

func AverageNumber(str string) float64 {

	var average float64
	if TestValidity(str) {
		splitted_array := strings.Split(str, "-")
		var numbers_count, numbers_total float64

		for index := 0; index < len(splitted_array); index++ {
			if modChecker(index) {
				intVar, _ := strconv.Atoi(splitted_array[index])

				numbers_count++
				numbers_total += float64(intVar)
			}
		}
		average = (numbers_total / numbers_count)
	} else {
		return 0.0
	}
	return average
}

func WholeStory(str string) string {

	var storyText []string
	if TestValidity(str) {

		splittedArray := strings.Split(str, "-")
		for index := 0; index < len(splittedArray); index++ {
			if !modChecker(index) {
				storyText = append(storyText, splittedArray[index])
			}
		}
	} else {
		return ""
	}
	return strings.Join(storyText, " ")
}

func modChecker(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}