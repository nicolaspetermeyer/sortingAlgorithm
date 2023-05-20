package main

import (
	"runtime"
	"time"
)

func insertionSort(g *Game) {
	if g.index < 1 || g.index >= len(g.data) {
		return
	}

	key := data[g.index]
	j := g.index - 1

	//visualization stuff
	g.value = int(data[g.index-1])
	g.i = int(data[g.index])
	if g.index < len(g.data)-1 {
		g.nextValue = int(data[g.index+1])
	}
	//end visualization stuff

	// as long as the next value is greater than the current value, swap them

	for j >= 0 && data[j] > key {
		data[j+1] = data[j]
		j--
		g.comparisons++
	}

	data[j+1] = key
	g.wg.Add(1)
	go func(index int) {
		defer g.wg.Done()

		frequency := 0.0
		if index < len(g.data)-1 {
			frequency = MapFreq(int(g.data[index+1]), len(g.data))
		} else {
			frequency = MapFreq(int(g.data[index]), len(g.data))
		}

		p := play(g.c, frequency, time.Duration(g.delay)*time.Millisecond, *channelCount, g.f)

		g.m.Lock()
		g.players = append(g.players, p)
		g.m.Unlock()
		Sleep(g.delay)
	}(g.index)
	Sleep(g.delay)

	g.wg.Wait()
	runtime.KeepAlive(g.players)

	if g.index == len(g.data)-1 {
		g.sorted = true
	}
	Sleep(g.delay)
	g.iterations = g.numSorted
}
