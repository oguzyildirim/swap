package main

import (
	"fmt"
)

func OptimalPageReplacement(pageReferences []int, frames int) int {
	faults := 0
	pageTable := make([]int, frames)

	for i := 0; i < frames; i++ {
		pageTable[i] = -1 // Initialize page table with empty slots
	}

	for i := 0; i < len(pageReferences); i++ {
		found := false
		for j := 0; j < frames; j++ {
			if pageTable[j] == pageReferences[i] {
				found = true
				break
			}
		}

		if !found {
			faults++
			if i == len(pageReferences)-1 {
				// Last page reference, no future references available
				pageTable[0] = pageReferences[i]
			} else {
				maxIndex := -1
				maxDistance := -1

				// Find the page that won't be used for the longest duration in the future
				for j := 0; j < frames; j++ {
					found := false
					for k := i + 1; k < len(pageReferences); k++ {
						if pageTable[j] == pageReferences[k] {
							found = true
							if k > maxDistance {
								maxDistance = k
								maxIndex = j
							}
							break
						}
					}

					if !found {
						maxIndex = j
						break
					}
				}

				// Replace the page with the longest future distance
				pageTable[maxIndex] = pageReferences[i]
			}
		}
	}

	return faults
}

func main() {
	pageReferences := []int{1, 2, 3, 4, 1, 2, 5, 1, 2, 3, 4, 5, 2, 2, 3, 5, 5, 1, 1}
	frames := 3

	pageFaults := OptimalPageReplacement(pageReferences, frames)
	fmt.Println("Number of page faults:", pageFaults)
}
