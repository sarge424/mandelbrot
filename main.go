package main

import (
	"fmt"
	"sarge424/mandelbrot/utils"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth, winHeight int = 800, 600
const maxIters int = 50

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("mandelbrot", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	t := time.Now()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		t = time.Now()

		arr := utils.MakeSlice(winWidth, winHeight)

		step := 4 / float64(winWidth)

		var beeper sync.WaitGroup
		beeper.Add(winHeight)

		for i := 0; i < winHeight; i++ {
			go utils.Mandelbrot(-1, 0.3-step*float64(i), 3, arr[i:i+1], maxIters, &beeper)
		}

		beeper.Wait()

		for i := 0; i < winHeight; i++ {
			for j := 0; j < winWidth; j++ {
				frac := 1 - float32(arr[i][j])/float32(maxIters)
				c := 255 * frac
				renderer.SetDrawColor(uint8(c*0.1), uint8(c*0.7), uint8(c*0.6), 255)

				renderer.DrawPoint(int32(j), int32(i))
			}
		}

		renderer.Present()
		fmt.Println(time.Since(t))
		//sdl.Delay(16)

	}
}
