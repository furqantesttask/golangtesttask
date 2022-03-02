package task

import (
	"strconv"
	"strings"
	"math"
	"math/rand"
	"sort"
	"time"
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

func StoryStats(str string) (string, string, float64, []string) {

		var shortestWord, longestWord string
		var averageWordLength            float64
		var averageWordLengthList        []string
		var storyText []string   
		wordCount := 0.0         
		accumulatedLength := 0.0

	if TestValidity(str) {

		splitted_array := strings.Split(str, "-")

		for index := 0; index < len(splitted_array); index++ {
			if !modChecker(index) {
				storyText = append(storyText, splitted_array[index])
				wordCount++
				accumulatedLength += float64(len(splitted_array[index]))
			}
		}

		averageWordLength = accumulatedLength / wordCount

		sort.Slice(storyText, func(i, j int) bool {
			return len(storyText[i]) < len(storyText[j])
		})

		shortestWord = storyText[0]
		longestWord = storyText[len(storyText)-1] 

		for _, word := range storyText {
			if len(word) == int(math.Ceil(averageWordLength)) || len(word) == int(math.Floor(averageWordLength)) {
				averageWordLengthList = append(averageWordLengthList, word)
			}
		}
	} else {
		var emptyString []string
		return "", "", 0.0, emptyString
	}
	return shortestWord, longestWord, averageWordLength, averageWordLengthList
}

func Generate(validityFlag bool) string {

	var storySequence string
	rand.Seed(time.Now().UnixNano())
	maxWordLen := 4294967295
	totalSequences := 1 + rand.Intn(10)
	var word, number string

	for i := 0; i < totalSequences; i++ {
		word = randSeq(1 + rand.Intn(10))

		number = strconv.Itoa(rand.Intn(maxWordLen))

		storySequence += number + "-" + word + "-"
	}

	if validityFlag == true {
		storySequence = storySequence[:len(storySequence)-1]
	}
	return storySequence
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func modChecker(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}