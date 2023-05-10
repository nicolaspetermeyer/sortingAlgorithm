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

const (
	WIDTH  = 1280
	HEIGTH = 720
)


type Game struct {
	
	data   			[]float64
	j, k   			int
	delay 			int;
	//sortAlgorithm 	string;
	sorted	 		bool;
}

type input struct {
	algorithm string
	delay    int
	//state   int
}

var (
	data       []float64
	barWidth   float64
	barHeight  float64
	barSpacing = 3.0
)

//---------------- Choose Algorithm  -----------------
func readAlgorithm(reader *bufio.Reader) string {
	pl("Choose algorithm: ")
	pl("1. Bubble sort")
	pl("2. Merge sort")
	pl("3. Heap sort")
	pl("4. Quick sort")

	//Read string until newline
	algorithm, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You chose: " + algorithm)
	}
	return strings.TrimSpace(algorithm)
}

//----------------- Read Dataset Size  -----------------
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

//----------------- Read Delay  -----------------
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

//----------------- Create Slice  -----------------
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

//----------------- Sleep  -----------------
func Sleep(n int) {
	time.Sleep(time.Duration(n) * time.Millisecond)
}

//----------------- Update Game -----------------
func (g *Game) Update(screen *ebiten.Image) error {
	if g.sorted {
		return nil
	}

	bubbleSort(g.data)
	return nil
} 
	
	




//----------------- Draw Game -----------------
func (g *Game) Draw(screen *ebiten.Image) {


	for i, num := range data {
		x := 10 + (barWidth+barSpacing)*float64(i)
		y := HEIGTH - num*barHeight

		var c color.RGBA
		if g.k < len(g.data)-1 {
			if i == g.j || i == g.j+1 {
				
			} else {
				c = color.RGBA{255, 255, 255, 255}
			}
		} else {
			c = color.RGBA{0, 255, 0, 255}
		}
		
		ebitenutil.DrawRect(screen, float64(x), float64(y), float64(barWidth), float64(num*barHeight), color.White)

	}
}
//----------------- Game Layout -----------------

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return WIDTH, HEIGTH
}


func main() {

	reader := bufio.NewReader(os.Stdin)
	var input input
	input.algorithm = readAlgorithm(reader)
	n := readCount(reader)
	data = createSlice(n)
	input.delay = readDelay(reader)

	barWidth = (float64(WIDTH) - 20 - (float64(n) * barSpacing)) / float64(n)
	barHeight = float64(HEIGTH) / float64(n)

	

	
	
	switch input.algorithm {
	case "1":
		bubbleSort(data)
		
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
		data:   data,
		k:      0,
		j:      0,
		delay:  input.delay,
		sorted: false,
		

	}

	//Start the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
