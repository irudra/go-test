package main

import "strconv"

func main() {
	print("I am working on test")
}

//I had estimated the the function as 15 minute however I took one hour in the implementing this.
func testValidity(test string) bool {
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
			return false
		} else if charV == separator {
			if startSequence {
				num, err := strconv.ParseUint(test[startIndex:i], 10, 64)
				if err != nil {
					return false
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
	return startSequence && startIndex != 0
}
