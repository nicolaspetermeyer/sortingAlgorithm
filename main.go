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
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var pl = fmt.Println

type Game struct {
	nums   []float64
	i, j   int
	sorted bool
}

const (
	width  = 1280
	height = 720
)

var (
	data       []float64
	barWidth   float64
	barHeight  float64
	barSpacing = 3.0
)

// ----------------- Update Game -----------------
func (g *Game) Update(*ebiten.Image) error {

	if g.sorted {
		return nil
	}
	(g.nums)
	return nil
}

// ----------------- Draw Game -----------------
func (g *Game) Draw(screen *ebiten.Image) {
	for i, num := range data {
		x := 10 + (barWidth+barSpacing)*float64(i)
		y := height - num*barHeight

		ebitenutil.DrawRect(screen, float64(x), float64(y), float64(barWidth), float64(num*barHeight), color.White)

	}
}

//----------------- Game Layout -----------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return width, height
}

// ----------------- Choose Algorithm  -----------------
func readAlgorithm(reader *bufio.Reader) string {
	pl("Choose algorithm: ")
	pl("1. Bubble sort")
	pl("2. Merge sort")
	pl("3. Heap sort")
	pl("4. Quick sort")

	// Read string until newline
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

func main() {

	reader := bufio.NewReader(os.Stdin)

	algo := readAlgorithm(reader)
	n := readCount(reader)
	data = createSlice(n)

	barWidth = (float64(width) - 20 - (float64(n) * barSpacing)) / float64(n)
	barHeight = float64(height) / float64(n)

	switch algo {
	case "1":
		pl(data)
		bubbleSort(data)
		// case "2":
		// 	pl(data)
		// 	mergeSort(data)

		// case "3":
		// 	pl(data)
		// 	minHeap := NewMinHeap(data)
		// 	minHeap.sort(len(data))
		// 	minHeap.print()
		// 	fmt.Scanln()
		// case "4":
		// 	pl("Quick sort")
	}

	// Set up the game window
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Shuffled Integers")
	game := &Game{
		nums:   data,
		i:      0,
		j:      0,
		sorted: false,
	}

	// Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
