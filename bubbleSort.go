package main

import "time"

func bubbleSort(data []float64) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				// swap
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
		time.Sleep(100 * time.Millisecond)
		return
	}
}
