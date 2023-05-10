package main

import "fmt"

func bubbleSortStep(g *Game, data []float64, delay int, i int, j int, k int) {
	if i >= len(data)-1 {
		g.sorted = true
		return
	}

	lastUnsorted := len(data) - 1 - g.numSorted
	if i >= lastUnsorted {
		g.numSorted++
		return
	}

	if data[i] > data[i+1] {
		//swap
		data[i], data[i+1] = data[i+1], data[i]

	}
	if i < len(data)-1 {
		g.j = int(data[i])
		g.k = int(data[i+1])
	}

	i++
	g.i = i
	Sleep(delay)
	if lastUnsorted == 1 {
		g.sorted = true
	}
	fmt.Println(g.sorted)
	fmt.Println(lastUnsorted)

}
