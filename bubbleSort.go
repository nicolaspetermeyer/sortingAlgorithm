package main

import (
	"runtime"
	"time"
)

func bubbleSortStep(g *Game, i int) {
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

	// current comparison
	if i < len(data)-1 {
		g.value = int(data[i])
		g.nextValue = int(data[i+1])
	}
	// current index iteration
	i++
	g.index = i

	// last comparison
	if lastUnsorted == 1 {
		g.sorted = true
	}
	g.comparisons++
	g.iterations = g.numSorted
}
