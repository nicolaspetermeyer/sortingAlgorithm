package main

import (
	"runtime"
	"time"
)

func selectionSort(g *Game, min_idx int) {
	startIndex := min_idx
	for j := startIndex + 1; j < len(g.data); j++ {
		if g.data[j] < g.data[min_idx] {
			g.value = min_idx

			min_idx = j

		}

		g.temp = min_idx

	}
	g.comparisons++
	g.wg.Add(1)
	go func(index int) {
		defer g.wg.Done()
		frequency := MapFreq(int(g.data[index]), len(g.data))

		p := play(g.c, frequency, time.Duration(g.delay)*time.Millisecond, *channelCount, g.f)
		g.m.Lock()
		g.players = append(g.players, p)
		g.m.Unlock()
		Sleep(g.delay)
	}(g.index)
	Sleep(g.delay)
	g.wg.Wait()
	runtime.KeepAlive(g.players)

	if g.numSorted == len(g.data)-1 {
		g.sorted = true
	}
	g.iterations = g.numSorted
}
