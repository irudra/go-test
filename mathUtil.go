package main

func calculateAverage(numbers []uint64) uint64 {
	average := numbers[0]
	for i := 1; i < len(numbers); i++ {
		u := average*uint64(i) + numbers[i]
		average = u / uint64(i+1)
	}
	return average
}
