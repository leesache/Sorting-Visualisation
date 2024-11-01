/*
MIT License

Copyright (c) 2024 Rushan Valimkhanov

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	screenWidth  = 20
	screenHeight = 20
	arraySize    = 15
)

func main() {
	start := time.Now()
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating screen: %v", err)
	}
	if err := screen.Init(); err != nil {
		log.Fatalf("Error initializing screen: %v", err)
	}

	screen.SetStyle(tcell.StyleDefault.Background(tcell.ColorHotPink).Foreground(tcell.ColorWhite))

	slice := sliceRandomInitialize(arraySize)

	exitChan := make(chan struct{})

	go func() {
		for {
			if ev := screen.PollEvent(); ev != nil {
				switch ev.(type) {
				case *tcell.EventKey:
					close(exitChan)
					return
				case *tcell.EventResize:
					screen.Sync()
				}
			}
		}
	}()

	var count uint64
	isSortedNow := false
	var elapsedTime time.Duration
	flag := 0
	for {
		if !isSortedNow {
			elapsedTime = time.Since(start)
			if flag == 150000 {
				visualizeArray(screen, slice, count, elapsedTime, isSortedNow)
				flag = 0
			}
			flag++
			if !isSorted(slice) {
				count++
				shuffleSlice(slice)
			} else {
				isSortedNow = true
				visualizeArray(screen, slice, count, elapsedTime, isSortedNow)
			}
		}

		time.Sleep(1 * time.Nanosecond)
		select {
		case <-exitChan:
			screen.Fini()
			os.Exit(0)
		default:
		}
	}
}

func sliceRandomInitialize(len int) []int {
	slice := make([]int, len)
	for i := 0; i < len; i++ {
		slice[i] = myRand(1, screenHeight)
	}
	return slice
}

func shuffleSlice(arr []int) {
	for i := range arr {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func myRand(min int, max int) int {
	return rand.Intn(max-min+1) + min
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func visualizeArray(screen tcell.Screen, arr []int, iteration uint64, elapsedTime time.Duration, isSorted bool) {
	screen.Clear()

	for i, value := range arr {
		for j := 0; j < value; j++ {
			if j >= screenHeight {
				break
			}
			style := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorWhite)
			screen.SetContent(i*4, screenHeight-j-1, '█', nil, style)
		}
	}

	iterationText := fmt.Sprintf("Iteration: %d", iteration)
	operationsPerSecond := float64(iteration) / elapsedTime.Seconds()
	speedText := fmt.Sprintf("Speed: %.2f ops/sec", operationsPerSecond)
	estimatedTime := fmt.Sprintf("Estimated time: %ds", factorial(int64(arraySize))/int64(operationsPerSecond))

	yPosition := screenHeight + 2
	timeX := 1
	iterationX := 1
	speedX := 1
	estimatedTimeX := 1

	timeStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite)

	for _, ch := range speedText {
		screen.SetContent(speedX, yPosition-1, ch, nil, timeStyle)
		speedX++
	}

	for _, ch := range iterationText {
		screen.SetContent(iterationX, yPosition, ch, nil, timeStyle)
		iterationX++
	}

	for _, ch := range elapsedTime.String() {
		screen.SetContent(timeX, yPosition+1, ch, nil, timeStyle)
		timeX++
	}
	for _, ch := range estimatedTime {
		screen.SetContent(estimatedTimeX, yPosition+2, ch, nil, timeStyle)
		estimatedTimeX++
	}
	screen.Show()
}

func factorial(n int64) int64 {
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
