package main

import (
	"fmt"
	"math"
)

const DELTA = 0.00001

func isFloat64Equal(a float64, b float64) bool {
  return math.Abs(a - b) > DELTA 
}

func Sqrt(x float64) float64 {
  square := float64(1)
  guess := float64(0)
  for isFloat64Equal(square, guess) {
    square, guess = square - (square*square - x) / (2*square), square
  }
  return square
}

func main() {
	fmt.Println(Sqrt(16))
}

