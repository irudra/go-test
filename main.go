package main

import (
	"errors"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	print("I am working on test")
}

//I had estimated the the function as 15 minute however I took one hour in the implementing this.
func testValidity(test string) bool {
	valid, _, _ := parseString(test)
	return valid
}

// Estimated time 15 minute, time taken 5 minute
func wholeStory(test string) (string, error) {
	valid, _, texts := parseString(test)
	if valid {
		return strings.Join(texts, " "), nil
	}
	return "", errors.New("format is not valid for the given string")
}

// Estimated time 15 minute, time taken 5 minute
func averageNumber(test string) (uint64, error) {
	valid, numbers, _ := parseString(test)
	if valid {
		average := calculateAverage(numbers)
		return average, nil
	}
	return 0, errors.New("format is not valid for the given string")
}

// estimated time half an hour complete 3/4 hour
func storyStats(test string) (string, string, int, []string) {
	valid, _, texts := parseString(test)
	if valid {
		average := float64(len(texts[0]))
		shortString := texts[0]
		longString := texts[0]
		for i := 1; i < len(texts); i++ {
			u := average*float64(i) + float64(len(texts[i]))
			average = u / float64(i+1)
			if len(shortString) > len(texts[i]) {
				shortString = texts[i]
			}
			if len(longString) < len(texts[i]) {
				longString = texts[i]
			}
		}
		var averageLengthString []string
		for i := 0; i < len(texts); i++ {
			if int(average) == len(texts[i]) {
				averageLengthString = append(averageLengthString, texts[i])
			}
		}
		return shortString, longString, int(average), averageLengthString
	}
	return "", "", 0, nil
}

func parseString(test string) (bool, []uint64, []string) {
	var integers []uint64
	var texts []string
	separator := "-"
	startIndex := 0
	startSequence := true
	for i := 0; i < len(test); i++ {
		charV := string(test[i])
		if len(test)-1 == i {
			texts = append(texts, test[startIndex:i+1])
			startSequence = true
		} else if charV == " " {
			return false, nil, nil
		} else if charV == separator {
			if startSequence {
				num, err := strconv.ParseUint(test[startIndex:i], 10, 64)
				if err != nil {
					return false, nil, nil
				}
				integers = append(integers, num)
				startSequence = false
				startIndex = i + 1
			} else {
				texts = append(texts, test[startIndex:i])
				startIndex = i + 1
				startSequence = true
			}
		}
	}
	return startSequence && startIndex != 0, integers, texts
}

func generateRandomPattern(valid bool) string {
	// make sure we have at-least one - in the generated string
	lengthTotal := rand.Intn(10) + 1
	var pattern strings.Builder
	if valid {
		for i := 0; i <= lengthTotal; i++ {
			charSetNumeric := "1234567890"
			// Make sure we have at-least one length for the valid numeric field
			length := rand.Intn(5) + 1
			if i != 0 {
				pattern.WriteString("-")
			}
			for j := 0; j <= length; j++ {
				random := rand.Intn(len(charSetNumeric))
				randomChar := charSetNumeric[random]
				pattern.WriteString(string(randomChar))
			}
			// Make sure we have at-least one length for the valid string field
			length = rand.Intn(10) + 1
			charSetString := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP123456789"
			pattern.WriteString("-")
			for k := 0; k <= length; k++ {
				random := rand.Intn(len(charSetString))
				randomChar := charSetString[random]
				pattern.WriteString(string(randomChar))
			}
		}
	} else {
		lengthTotal := rand.Intn(20) + 1
		charSet := "abcdedfghijklmnopqrst ABCDEFGHIJKLMNOP123456789"
		for i := 0; i < lengthTotal; i++ {
			random := rand.Intn(len(charSet))
			randomChar := charSet[random]
			pattern.WriteString(string(randomChar))
		}
	}
	return pattern.String()
}
