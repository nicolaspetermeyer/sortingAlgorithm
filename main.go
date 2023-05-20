package main

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/hajimehoshi/oto/v2"
)

var pl = fmt.Println

const (
	WIDTH  = 1280
	HEIGTH = 730
)

type Game struct {
	data                             []float64
	temp, i, index, value, nextValue int
	delay                            int
	sorted                           bool
	numSorted                        int
	algorithm                        string

	// sort infos
	iterations   int
	swaps        int
	comparisons  int
	array_access int

	// sound
	players []oto.Player
	f       int
	c       *oto.Context
	wg      sync.WaitGroup
	m       sync.Mutex
}

var (
	data       []float64
	barWidth   float64
	barHeight  float64
	barSpacing = 3.0
	antialias  = true
	algorithm  string
	delay      int
	co         color.RGBA
)

func MapFreq(num, n int) float64 {
	return float64(num*2000/n + 100)
}

// ---------------- Choose Algorithm  -----------------
func readAlgorithm(reader *bufio.Reader) string {
	pl("Choose algorithm: ")
	pl("1. Bubble sort")
	pl("2. Insertion sort")
	pl("3. Selection sort")

	//Read string until newline
	algorithm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You chose: " + algorithm)
	}
	return strings.TrimSpace(algorithm)
}

// ----------------- Read Dataset Size  -----------------
func readCount(reader *bufio.Reader) int {
	pl("Enter n: ")

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	size, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You entered: " + n)
	}
	return size
}

// ----------------- Read Delay  -----------------
func readDelay(reader *bufio.Reader) int {
	pl("Enter delay in ms: ")

	n, _ := reader.ReadString('\n')
	n = strings.TrimSpace(n)
	delay, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You entered: " + n)
	}
	return delay
}

// ----------------- Create Slice  -----------------
func createSlice(size int) []float64 {
	data := make([]float64, size)
	for i := range data {
		data[i] = float64((i + 1))
	}
	//shuffle slice
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	pl(data)
	return data
}

// ----------------- Sleep  -----------------
func Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

// ----------------- Update Game -----------------
func (g *Game) Update() error {

	if g.sorted {
		return nil
	}

	if g.algorithm == "1" {
		algorithm = "Bubble sort"
		bubbleSortStep(g, g.index)
		if g.index >= len(data)-1-g.numSorted {
			g.index = 0
			g.numSorted++
		}
	} else if g.algorithm == "2" {
		algorithm = "Insertion sort"

		if !g.sorted {
			insertionSort(g)
		}
		if g.index < len(g.data) {
			g.index++
		}

	} else if g.algorithm == "3" {
		algorithm = "Selection sort"
		if !g.sorted {
			selectionSort(g, g.numSorted)

			if g.index < len(g.data)-1 {
				g.index++
			} else {
				g.data[g.numSorted], g.data[g.temp] = g.data[g.temp], g.data[g.numSorted]
				g.index = g.numSorted
				g.numSorted++
			}

		}
	}
	return nil
}

// ----------------- Draw Game -----------------
func (g *Game) Draw(screen *ebiten.Image) {

	for i, num := range g.data {
		x := 10 + (barWidth+barSpacing)*float64(i)
		y := HEIGTH - num*barHeight

		if g.algorithm == "1" {
			if num == float64(g.value) {
				co = color.RGBA{0xff, 0x00, 0x00, 0xff} //red
			} else if num == float64(g.nextValue) {
				co = color.RGBA{0x00, 0x00, 0xff, 0xff} //blue
			} else if i >= len(g.data)-g.numSorted {
				co = color.RGBA{0x00, 0xff, 0xff, 0xff} //light blue
			} else {
				co = color.RGBA{0xff, 0xff, 0xff, 0xff} //white
			}
		}
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), co, antialias)

		if g.algorithm == "2" {
			if num == float64(g.value) {
				co = color.RGBA{0xff, 0x00, 0x00, 0xff} //red
			} else if num == float64(g.nextValue) {
				co = color.RGBA{0x00, 0x00, 0xff, 0xff} //blue
			} else if num == float64(g.i) {
				co = color.RGBA{0x00, 0xff, 0x00, 0xff} //green
			} else if num == float64(g.temp) {
				co = color.RGBA{0x80, 0x00, 0x00, 0xff} //dark red
			} else {
				co = color.RGBA{0xff, 0xff, 0xff, 0xff} //white
			}
		}
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), co, antialias)

		if g.algorithm == "3" {
			if i == g.index {
				co = color.RGBA{0x00, 0x00, 0xff, 0xff} //blue
			} else if num == float64(g.data[g.temp]) {
				co = color.RGBA{0x80, 0x00, 0x00, 0xff} //red
			} else if i < g.numSorted {
				co = color.RGBA{0x00, 0xff, 0xff, 0xff} //light blue
			} else {
				co = color.RGBA{0xff, 0xff, 0xff, 0xff} //white
			}
		}
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), co, antialias)
	}
	if g.sorted {
		for i, num := range g.data {
			x := 10 + (barWidth+barSpacing)*float64(i)
			y := HEIGTH - num*barHeight
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0xff, 0x00, 0xff}, antialias)
		}
	}
	msg := fmt.Sprint(algorithm, " - ", g.iterations, " Iterations, ", g.comparisons, " Comparisons, ", g.delay, " ms delay")
	ebitenutil.DebugPrintAt(screen, msg, 10, 10)
}

// ----------------- Game Layout -----------------
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGTH
}

func (g *Game) newContext() error {

	switch *format {
	case "f32le":
		g.f = oto.FormatFloat32LE
	case "u8":
		g.f = oto.FormatUnsignedInt8
	case "s16le":
		g.f = oto.FormatSignedInt16LE
	default:
		return fmt.Errorf("format must be u8, s16le, or f32le but: %s", *format)
	}

	c, ready, err := oto.NewContext(*sampleRate, *channelCount, g.f)
	if err != nil {
		return err
	}

	g.c = c
	<-ready
	return nil
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	algorithm = readAlgorithm(reader)
	n := readCount(reader)
	delay = readDelay(reader)
	data = createSlice(n)
	barWidth = (float64(WIDTH) - 20 - (float64(n) * barSpacing)) / float64(n)
	barHeight = float64(HEIGTH-10) / float64(n)

	//Set up the game window
	ebiten.SetWindowSize(WIDTH, HEIGTH)
	ebiten.SetWindowTitle("Shuffled Integers")
	game := &Game{
		data:         data,
		index:        0,
		value:        0,
		nextValue:    0,
		i:            1,
		temp:         0,
		delay:        delay,
		sorted:       false,
		numSorted:    0,
		algorithm:    algorithm,
		swaps:        0,
		comparisons:  0,
		array_access: 0,
		iterations:   0,
	}

	go game.newContext()

	//Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
