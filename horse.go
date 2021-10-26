package main

import "fmt"

func main() {
	fmt.Println(racesForFastestTwo(30))
}

func racesForFastestTwo(horses int) int {
	races, candidates := 0, horses
	// If there are more than two horses
	// continue to race
	for candidates > 2 {
		// Full races at capacity 5
		numOfFullRaces, remainder := candidates/5, candidates%5
		races += numOfFullRaces
		// If there is a remainder > 2 horses
		// we must race them and if less we don't need to
		if remainder > 2 {
			races++
		}
		// For each race of 5 we take the fastest two
		// The fastest two of all 30 horses will come out
		// Top two in any race of 5
		candidates = 2 * (numOfFullRaces)
		// If the remainder is less than 3
		// We would not need to race them
		// Otherwise we would and that will give us two candidates
		if remainder < 3 {
			candidates += remainder
		} else {
			candidates += 2
		}
	}
	return races
}
