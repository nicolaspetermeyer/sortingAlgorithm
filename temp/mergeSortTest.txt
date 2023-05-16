package main

import (
	"fmt"
)

func mergeSort(g *Game, data []float64) []float64 {
	if len(data) <= 1 {
		return data
	}
	first := mergeSort(g, data[:len(data)/2])
	second := mergeSort(g, data[len(data)/2:])

	fmt.Println("test: ", g.tmp)
	return merge(g, first, second)

}

func merge(g *Game, a, b []float64) []float64 {
	i, j := 0, 0
	final := []float64{}
	g.left = a
	g.right = b
	fmt.Println(a, b)
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
		fmt.Println("first print: ", final)

	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
		fmt.Println("second print: ", final)
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
		fmt.Println("third print: ", final)
	}
	pl("last print: ", final)

	g.tmp = final
	fmt.Println("test 2: ", g.tmp)
	return final
}
