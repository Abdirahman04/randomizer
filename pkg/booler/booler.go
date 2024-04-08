package booler

import "math/rand"

func GetBool() bool {
  bools := [2]bool{true, false}
  return bools[rand.Intn(2)]
}
