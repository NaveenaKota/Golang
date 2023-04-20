package calculator

import "math"

func Add(x,y,z int) int{
	return x + y + z
}

func Diff(x, y int) int {
	return x - y
}

func Mult(x, y int) int {
	return x * y
}

func Div(x, y int) int {
	return x/y
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Pow(x,y float64) float64 {
	return math.Pow(x,y)
}

func Pow10(n int) float64 {
	return math.Pow10(n)
}

func Sincos(x float64) (sin, cos float64) {
	return math.Sin(x), math.Cos(x)
}
func Exp(x float64) float64 {
	return math.Exp(x)
}

func Log(x float64) float64 {
	return math.Log2(x)
}

