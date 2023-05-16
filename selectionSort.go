package main

func selectionSortStep(g *Game) {
	if g.index+1 < 1 || g.index+1 >= len(g.data) {
		return
	}

	// as we have an array from 1 to n, the index g.index represents the smallest element
	min_idx := g.index + 1
	for j := g.index; j < len(g.data); j++ {
		if g.data[j] < g.data[min_idx] {
			min_idx = j
		}
	}
	// swap the minimum element with the first element
	g.data[g.index], g.data[min_idx] = g.data[min_idx], g.data[g.index]

	// current index iteration
	g.i = g.index
	if g.index < len(g.data)-1 {
		g.value = int(g.data[g.index])
		g.nextValue = int(g.data[g.index+1])
	}

	if g.index == len(g.data)-1 {
		g.sorted = true
	}
	Sleep(delay)

}
