package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var pl = fmt.Println

const (
	WIDTH  = 1280
	HEIGTH = 720
)

// type bar struct {
// 	rect  pixel.Rect
// 	color color.Color
// }

// type bar struct {
// 	rect  *ebiten.Image
// 	color color.RGBA
// }

type Game struct {
	data                       []float64
	i, index, value, nextValue int
	delay                      int
	sorted                     bool
	numSorted                  int
	algorithm                  string

	// sort infos
	iterations   int
	swaps        int
	comparisons  int
	array_access int

	// insertion sort
}

var (
	data       []float64
	barWidth   float64
	barHeight  float64
	barSpacing = 3.0
	antialias  = true
	algorithm  string
	delay      int
)

// // ---------------- Choose Algorithm  -----------------
// func readAlgorithm(reader *bufio.Reader) string {
// 	pl("Choose algorithm: ")
// 	pl("1. Bubble sort")
// 	pl("2. Merge sort")
// 	pl("3. Heap sort")
// 	pl("4. Quick sort")

// 	//Read string until newline
// 	algorithm, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("You chose: " + algorithm)
// 	}
// 	return strings.TrimSpace(algorithm)
// }

// // ----------------- Read Dataset Size  -----------------
// func readCount(reader *bufio.Reader) int {
// 	pl("Enter n: ")

// 	n, _ := reader.ReadString('\n')
// 	n = strings.TrimSpace(n)
// 	size, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("You entered: " + n)
// 	}
// 	return size
// }

// // ----------------- Read Delay  -----------------
// func readDelay(reader *bufio.Reader) int {
// 	pl("Enter delay in ms: ")

// 	n, _ := reader.ReadString('\n')
// 	n = strings.TrimSpace(n)
// 	delay, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("You entered: " + n)
// 	}
// 	return delay
// }

// ----------------- Create Slice  -----------------
func createSlice(size int) []float64 {
	data := make([]float64, size)
	for i := 0; i < len(data); i++ {
		data[i] = float64( /*height / size */ (i + 1))
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
		bubbleSortStep(g, g.index)
		if g.index >= len(data)-1-g.numSorted {
			g.index = 0
			g.numSorted++
		}
	} else if g.algorithm == "2" {
		if !g.sorted {
			insertionSort(g)
			if g.index < len(g.data) {
				g.index++
			}
		}
	} else if g.algorithm == "3" {
		if !g.sorted {
			selectionSortStep(g)
			if g.index < len(g.data)-1 {
				g.index++
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

		if num == float64(g.value) {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0xff, 0x00, 0x00, 0xff}, antialias)
		} else if num == float64(g.nextValue) {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0x00, 0xff, 0xff}, antialias)
		} else if num == float64(g.i) && g.algorithm == "2" {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0xff, 0x00, 0xff}, antialias)
		} else if g.algorithm == "3" && num == float64(g.index) {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0xff, 0x00, 0xff}, antialias)

		} else {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.White, antialias)
		}
	}
	if g.sorted {
		for i, num := range g.data {
			x := 10 + (barWidth+barSpacing)*float64(i)
			y := HEIGTH - num*barHeight
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0xff, 0x00, 0xff}, antialias)
		}
	}

}

// //----------------- Game Layout -----------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGTH
}

func main() {

	//reader := bufio.NewReader(os.Stdin)

	// algorithm = readAlgorithm(reader)
	// n := readCount(reader)
	// data = createSlice(n)
	// delay = readDelay(reader)

	algorithm = "3"
	n := 50
	data = createSlice(n)
	delay = 100

	barWidth = (float64(WIDTH) - 20 - (float64(n) * barSpacing)) / float64(n)
	barHeight = float64(HEIGTH) / float64(n)

	//Set up the game window
	ebiten.SetWindowSize(WIDTH, HEIGTH)
	ebiten.SetWindowTitle("Shuffled Integers")
	game := &Game{
		data:         data,
		index:        0,
		value:        0,
		nextValue:    0,
		i:            1,
		delay:        delay,
		sorted:       false,
		numSorted:    0,
		algorithm:    algorithm,
		swaps:        0,
		comparisons:  0,
		array_access: 0,
		iterations:   0,
	}

	//Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
