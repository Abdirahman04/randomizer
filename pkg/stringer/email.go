package stringer

const EmailSuffix = "@gmail.com"

func RandomEmail(ln int) string {
  return RandomStringLower(ln) + EmailSuffix
}
