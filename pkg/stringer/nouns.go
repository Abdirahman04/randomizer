package stringer

import (
	"math/rand"
	"strings"
)

func GetSyllable() string {
  syl := [27]string{"b","c","ch","d","f","g","gh","h","j","k","kh","l","m","n","p","ph","q","r","s","sh","t","th","v","w","x","y","z"}
  vow := [5]string{"a","e","i","o","u"}
  return syl[rand.Intn(27)] + vow[rand.Intn(5)]
}

func GetEndConsonant() string {
  con := [15]string{"b","c","d","f","g","j","k","l","m","n","p","s","t","y","z"}
  return con[rand.Intn(15)]
}

func GetVowel() string {
  vow := [5]string{"a","e","i","o","u"}
  return vow[rand.Intn(5)]
}

func RandomNounLower(ln int) string {
  if ln == 1 {
    return GetSyllable()
  }
  str := GetSyllable()
  if ln > 2 {
    for i := 0;i < ln-2;i++ {
      str += GetSyllable()
    }
  }
  return str + GetEndConsonant()
}

func RandomNounUpper(ln int) string {
  return strings.ToUpper(RandomNounLower(ln))
}

func RandomNounCap(ln int) string {
  return strings.ToUpper(GetEndConsonant()) + GetVowel() + RandomNounLower(ln-1)
}
