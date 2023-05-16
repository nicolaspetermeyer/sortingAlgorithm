package main

import "fmt"

func bubbleSortStep(g *Game, i int) {
	if i >= len(data)-1 {
		g.sorted = true
		return
	}

	// remove sorted elements from the end
	lastUnsorted := len(data) - 1 - g.numSorted
	if i >= lastUnsorted {
		g.numSorted++
		return
	}

	// swap
	if data[i] > data[i+1] {
		data[i], data[i+1] = data[i+1], data[i]
		g.swaps++
	}
	// current comparison
	if i < len(data)-1 {
		g.value = int(data[i])
		g.nextValue = int(data[i+1])
	}
	// current index iteration
	i++
	g.index = i
	Sleep(delay)

	// last comparison
	if lastUnsorted == 1 {
		g.sorted = true
	}
	fmt.Println("swap: ", g.swaps)
	g.comparisons++
	fmt.Println("comparison: ", g.comparisons)
}
