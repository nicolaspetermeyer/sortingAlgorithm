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
// 	rect  *ebiten.Image
// 	color color.RGBA
// }

type Game struct {
	data      []float64
	i, j, k   int
	delay     int
	sorted    bool
	numSorted int
}

type input struct {
	algorithm string
	delay     int
}

var (
	data       []float64
	barWidth   float64
	barHeight  float64
	barSpacing = 3.0
	antialias  = true
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
	time.Sleep(time.Duration(n) * time.Microsecond)
}

// ----------------- Update Game -----------------
func (g *Game) Update() error {

	if g.sorted {
		return nil
	}

	bubbleSortStep(g, g.data, g.delay, g.i, g.j, g.k)
	if g.i >= len(data)-1-g.numSorted {
		g.i = 0
		g.numSorted++
	}

	return nil
}

// ----------------- Draw Game -----------------
func (g *Game) Draw(screen *ebiten.Image) {

	for i, num := range g.data {
		x := 10 + (barWidth+barSpacing)*float64(i)
		y := HEIGTH - num*barHeight

		if num == float64(g.j) {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0xff, 0x00, 0x00, 0xff}, antialias)
		} else if num == float64(g.k) {
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(barWidth), float32(num*barHeight), color.RGBA{0x00, 0x00, 0xff, 0xff}, antialias)
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

//----------------- Game Layout -----------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGTH
}

func main() {

	//reader := bufio.NewReader(os.Stdin)
	var input input
	// input.algorithm = readAlgorithm(reader)
	// n := readCount(reader)
	// data = createSlice(n)
	// input.delay = readDelay(reader)

	input.algorithm = "1"
	n := 20
	data = createSlice(n)
	input.delay = 1

	barWidth = (float64(WIDTH) - 20 - (float64(n) * barSpacing)) / float64(n)
	barHeight = float64(HEIGTH) / float64(n)

	switch input.algorithm {
	case "1":
	//bubbleSort(data, input.delay, progress)

	// case "2":
	// 	data = mergeSort(data)

	// 	case "3":
	// 		pl(data)
	// 		minHeap := NewMinHeap(data)
	// 		minHeap.sort(len(data))
	// 		minHeap.print()
	// 		fmt.Scanln()
	// 	case "4":
	// 		pl("Quick sort")
	default:
		pl("Invalid input")
		return
	}

	//Set up the game window
	ebiten.SetWindowSize(WIDTH, HEIGTH)
	ebiten.SetWindowTitle("Shuffled Integers")
	game := &Game{
		data:      data,
		i:         0,
		j:         0,
		k:         0,
		delay:     input.delay,
		sorted:    false,
		numSorted: 0,
	}

	//Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
