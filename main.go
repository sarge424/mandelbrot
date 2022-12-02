package main

import (
	"fmt"
	"sarge424/mandelbrot/utils"
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

		arr := utils.MakeSlice(winWidth, winHeight)

		s := winHeight / 3
		step := 4 / float64(winWidth) * float64(s)

		utils.Mandelbrot(-3, 1, 4, arr[0:s], maxIters)
		utils.Mandelbrot(-3, 1-step, 4, arr[s:2*s], maxIters)
		utils.Mandelbrot(-3, 1-2*step, 4, arr[2*s:3*s], maxIters)

		for i := 0; i < winHeight; i++ {
			for j := 0; j < winWidth; j++ {
				frac := 1 - float32(arr[i][j])/float32(maxIters)
				c := 255 * frac

				renderer.SetDrawColor(uint8(c*0.1), uint8(c*0.7), uint8(c*0.6), 255)
				renderer.DrawPoint(int32(j), int32(i))
			}
		}

		fmt.Println(time.Since(t))
		t = time.Now()

		renderer.Present()
		sdl.Delay(16)

	}
}
