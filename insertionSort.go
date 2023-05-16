package main

import "fmt"

func insertionSort(g *Game) {
	if g.current < 1 || g.current >= len(g.data) {
		return
	}

	key := data[g.current]
	j := g.current - 1

	g.j = int(data[g.current-1])
	g.i = int(data[g.current])
	if g.current < len(g.data)-1 {
		g.k = int(data[g.current+1])
	}
	for j >= 0 && data[j] > key {
		data[j+1] = data[j]
		j--
	}
	data[j+1] = key

	if g.current == len(g.data)-1 {
		g.sorted = true
	}
	//g.j = int(data[int(key)-1])
	fmt.Println(g.current, g.j, g.k, key, j)
	Sleep(delay)

}

// notes

// if not, swap data[1] and data[0] and return data
// Is data [2] sorted?
// if not, swap data[2] and data[1]
