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

func modChecker(i int) bool {
	if i%2 == 0 {
		return true
	} else {
		return false
	}
}
