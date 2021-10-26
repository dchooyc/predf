package main

import "fmt"

func main() {
	array := genTestArray()
	fmt.Println(annualDinner(array))
}

func annualDinner(array []int) int {
	dict, offset := make(map[int]int), 888*400
	for i := 0; i < len(array); i++ {
		// There will be 400 pairs that add up to 888
		// We take their total 888 * 400
		// And subtract each element in the array
		// Resulting in a negative of the number without a pair
		offset -= array[i]
		// This makes a key of the number plate with it's value
		// As it's index in the array
		dict[array[i]] = i
	}
	// We will have a negative number here
	// Because of the last number without a complement
	offset *= -1
	return dict[offset]
}

// Generates an array with numbers 1-400 and it's complement
// Then inserts 405 a number without a complement at index 200
func genTestArray() []int {
	array, num := []int{}, 1
	for i := 0; i < 400; i++ {
		complement := 888 - num
		array = append(array, complement)
		array = append(array, num)
		num++
	}
	array = append(array[:201], array[200:]...)
	array[200] = 405
	return array
}
