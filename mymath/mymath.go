package mymath

import "math"

func MySqrt(x float64) float64 {
	return math.Sqrt(x)
}

func MyCeil(x float64) float64 {
	return math.Ceil(x)
}

func MyFloor(x float64) float64 {
	return math.Floor(x)
}

func MyPow(x, y float64) float64 {
	return math.Pow(x, y)
}

func MyMax(x, y float64) float64 {
	return math.Max(x, y)
}

func MyMin(x, y float64) float64 {
	return math.Min(x, y)
}
