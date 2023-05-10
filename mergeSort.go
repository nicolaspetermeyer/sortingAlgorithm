package main

import "time"

func mergeSortStep(g *Game, data []float64, delay int, start int, end int, mid int) {
	left := make([]float64, mid-start+1)
	right := make([]float64, end-mid)

	for i := 0; i < len(left); i++ {
		left[i] = data[start+i]
	}

	for j := 0; j < len(right); j++ {
		right[j] = data[mid+1+j]
	}

	i := 0
	j := 0
	k := start

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			data[k] = left[i]
			i++
		} else {
			data[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		data[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		data[k] = right[j]
		j++
		k++
	}

	g.start = start
	g.end = end
	g.mid = mid

	time.Sleep(time.Duration(delay) * time.Millisecond)
}
