package main

func selectionSort(g *Game, min_idx int) {
	startIndex := min_idx
	for j := startIndex + 1; j < len(g.data); j++ {
		if g.data[j] < g.data[min_idx] {
			g.value = min_idx

			min_idx = j
		}
		g.temp = min_idx

	}
	if g.numSorted == len(g.data)-1 {
		g.sorted = true
	}
	Sleep(delay)
}
