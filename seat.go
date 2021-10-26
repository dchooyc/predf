package main

import (
	"fmt"
	"strconv"
)

func main() {
	// An array where most popular seat is 1
	// and there were four sold
	array := []int{1, 1, 1, 1, 2, 2, 3, 3}
	seatNumber, numSold := popSeat(array)
	fmt.Println("The most popular seat number is: " + strconv.Itoa(seatNumber))
	fmt.Println("The number of tickets sold for this seat number is: " + strconv.Itoa(numSold))
}

func popSeat(array []int) (seatNumber, numSold int) {
	dict, maxNumSold, mostPopSeat := map[int]int{array[0]: 1}, 1, array[0]
	for i := 1; i < len(array); i++ {
		// Check if we have the number in the dictionary
		if val, ok := dict[array[i]]; ok {
			// Since we found another entry we add one
			// To it's count
			newVal := val + 1
			dict[array[i]] = newVal
			// If it's new count is higher than the highest
			// count so far, we assign a new maxNumSold so far
			// we also change the mostPopSeat to the current seat
			if newVal > maxNumSold {
				maxNumSold = newVal
				mostPopSeat = array[i]
			}
		} else {
			// Create the seat number in dictionary
			// Give it a count of one
			dict[array[i]] = 1
		}
	}
	return mostPopSeat, maxNumSold
}
