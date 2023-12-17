package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFactorial(t *testing.T) {
	result := Factorial(10)

	if result != 3628800 {
		t.Errorf("Factorial(10) = %d; want 3628800", result)
	}
}

func TestHelloWorldFunc(t *testing.T) {
	var buf bytes.Buffer

	fmt.Fprint(&buf, HelloWorld())

	output := buf.String()
	expected := "Hello world!"

	if output != expected {
		t.Errorf("Unexpected output: %s", output)
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, test := range tests {
		result := Fibonacci(test.input)
		if result != test.expected {
			t.Errorf("Fibonacci(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

func average(xs []float64) float64 {
	sum := 0.0
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func generateSlice(size int) []float64 {
	rand.Seed(time.Now().UnixNano())
	slice := make([]float64, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Float64() * 100
	}
	return slice
}

func TestAverage(t *testing.T) {
	tests := []struct {
		input    []float64
		expected float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3},
		{[]float64{10, 20, 30, 40, 50}, 30},
		{[]float64{-1, -2, -3, -4, -5}, -3},
	}

	for _, test := range tests {
		result := average(test.input)
		if result != test.expected {
			t.Errorf("average(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}

	for i := 0; i < 10; i++ {
		size := rand.Intn(100)
		slice := generateSlice(size)
		expected := average(slice)
		result := average(slice)
		if result != expected {
			t.Errorf("average(%v) = %v; expected %v", slice, result, expected)
		}
	}
}

func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = fibonacciFormula(7)
    }
}