package stringer

import (
	"math/rand"
	"strings"
)

func GetChar() string {
  chars := [26]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

  return chars[rand.Intn(26)]
}

func RandomStringLower(ln int) string {
  var str string
  for i := 0;i < ln;i++ {
    str += GetChar()
  }
  return str
}

func RandomStringUpper(ln int) string {
  return strings.ToUpper(RandomStringLower(ln))
}

func RandomStringCap(ln int) string {
  str := strings.ToUpper(GetChar())
  str += RandomStringLower(ln-1)
  return str
}
