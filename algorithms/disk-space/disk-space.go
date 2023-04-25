package main

import "fmt"

func ChunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func Segment(x int32, space []int32) (int32, []int32) {
	// Write your code here

	min := make([]int32, 0)
	val := space[0]

	// get segments min
	for i := 1; i < len(space); i++ {
		for j := 0; j < int(x); j++ {
			if val > space[j] {
				val = space[j]
			}
		}
		val := space[i]
		min = append(min, val)
		space = space[i:]

		i = 0
	}

	min = append(min, val)

	// get max
	val = 0
	for _, n := range min {
		if val < n {
			val = n
		}
	}

	return val, min
}

// add to map and decrese slice one by one
func SliceToMap(x int, si []int) map[int][]int {

	m := make(map[int][]int)

	fmt.Println(si)
	j := 0
	for i := 1; i < len(si); i++ {
		m[j] = si
		si = si[i:]
		fmt.Println(si)
		i = 0
		j++
	}

	m[j] = si

	return m
}

func RemoveItemsFromSlice(si []int) []int {
	fmt.Println(si)
	for i := 1; i < len(si); i++ {
		si = si[i:]
		fmt.Println(si)
		i = 0
	}
	return si
}
