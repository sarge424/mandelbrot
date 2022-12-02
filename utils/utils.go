package utils

func mag2(c complex128) float64 {
	return real(c)*real(c) + imag(c)*imag(c)
}

func itersToInfinity(c complex128, maxIters int) int {
	var z complex128 = 0 + 0i

	var i int = 0
	for ; i < maxIters && mag2(z) < 16; i++ {
		z = z*z + c
	}

	return i
}

func makeSlice(w, h int) [][]int {
	a := [][]int{}

	for i := 0; i < h; i++ {
		r := []int{}
		for j := 0; j < w; j++ {
			r = append(r, 0)
		}
		a = append(a, r)
	}

	return a
}

func Mandelbrot(x, y, scale float64, w, h, iters int) [][]int {
	arr := makeSlice(w, h)

	var r, c int
	step := scale / float64(w)
	for i := y; r < h; i -= step {
		c = 0
		for j := float64(x); c < w; j += step {
			z := complex(j, i)
			arr[r][c] = itersToInfinity(z, iters)
			c++
		}
		r++
	}

	return arr
}
