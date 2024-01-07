package mymath

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	x := 4.0
	expected := 2.0
	result := MySqrt(x)
	if result != expected {
		t.Errorf("MySqrt(%v) = %v; expected %v", x, result, expected)
	}
}

func TestMyCeil(t *testing.T) {
	x := 3.5
	expected := 4.0
	result := MyCeil(x)

	if result != expected {
		t.Errorf("MyCeil(%v) = %v; expected %v", x, result, expected)
	}
}

func TestMyFloor(t *testing.T) {
	x := 3.5
	expected := 3.0
	result := MyFloor(x)
	if result != expected {
		t.Errorf("MyFloor(%v) = %v; expected %v", x, result, expected)
	}
}

func TestMyPow(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 8.0
	result := MyPow(x, y)
	if result != expected {
		t.Errorf("MyPow(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}

func TestMyMax(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 3.0
	result := MyMax(x, y)

	if result != expected {
		t.Errorf("MyMax(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}

func TestMyMin(t *testing.T) {
	x := 2.0
	y := 3.0
	expected := 2.0
	result := MyMin(x, y)

	if result != expected {
		t.Errorf("MyMin(%v, %v) = %v; expected %v", x, y, result, expected)
	}
}
