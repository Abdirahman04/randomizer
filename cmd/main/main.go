package main

import (
	"fmt"

	"github.com/Abdirahman04/randomizer/pkg/stringer"
)

func main() {
  fmt.Println("<<<<<     RANDOMIZER     >>>>>")

  fmt.Println("Lower => ", stringer.RandomStringLower(10))
  fmt.Println("Upper => ", stringer.RandomStringUpper(10))
  fmt.Println("Capitalized => ", stringer.RandomStringCap(10))

  fmt.Println("Email => ", stringer.RandomEmail(10))
}
