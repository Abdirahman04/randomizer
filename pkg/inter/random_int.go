package inter

import (
	"math"
	"math/rand"
)

func RandomNumber(ln int) int {
  x := math.Pow(10, float64(ln))
  return rand.Intn(int(x))
}
