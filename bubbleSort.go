package main

func bubbleSortStep(g *Game, data []float64, delay int, i int, j int, k int) {
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

	}
	// current comparison
	if i < len(data)-1 {
		g.j = int(data[i])
		g.k = int(data[i+1])
	}
	// current index iteration
	i++
	g.i = i
	Sleep(delay)

	// last comparison
	if lastUnsorted == 1 {
		g.sorted = true
	}
}
