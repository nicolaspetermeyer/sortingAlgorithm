package main

func insertionSort(g *Game) {
	if g.index < 1 || g.index >= len(g.data) {
		return
	}

	key := data[g.index]
	j := g.index - 1

	g.value = int(data[g.index-1])
	g.i = int(data[g.index])
	if g.index < len(g.data)-1 {
		g.nextValue = int(data[g.index+1])
	}
	for j >= 0 && data[j] > key {
		data[j+1] = data[j]
		j--

	}
	data[j+1] = key

	if g.index == len(g.data)-1 {
		g.sorted = true
	}
	Sleep(delay)
}
