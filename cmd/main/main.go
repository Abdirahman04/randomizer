package main

import (
	"fmt"

	"github.com/Abdirahman04/randomizer/pkg/booler"
	"github.com/Abdirahman04/randomizer/pkg/inter"
	"github.com/Abdirahman04/randomizer/pkg/stringer"
)

func main() {
  fmt.Println("<<<<<     RANDOMIZER     >>>>>")

  fmt.Println("Character => ", stringer.GetChar())
  fmt.Println("Lower => ", stringer.RandomStringLower(10))
  fmt.Println("Upper => ", stringer.RandomStringUpper(10))
  fmt.Println("Capitalized => ", stringer.RandomStringCap(10))

  fmt.Println("Email => ", stringer.RandomEmail(10))
  fmt.Println("Syllable => ", stringer.GetSyllable())
  fmt.Println("Noun Email => ", stringer.NounEmail(4))

  fmt.Println("Consonant => ", stringer.GetEndConsonant())
  fmt.Println("Vowel => ", stringer.GetVowel())

  fmt.Println("Noun Lower => ", stringer.RandomNounLower(4))
  fmt.Println("Noun Upper => ", stringer.RandomNounUpper(4))
  fmt.Println("Noun Capitalized => ", stringer.RandomNounCap(4))

  fmt.Println("\nBoolean => ", booler.GetBool())

  fmt.Println("\nNumber => ", inter.RandomNumber(6))
}
