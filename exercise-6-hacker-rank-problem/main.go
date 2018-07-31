package main

import (
  "fmt"
)

func main() {
  // testing camelcase
  count := camelcase("testPhrase")
  fmt.Println(count)

  //testing caesar cipher
  ciphered := caesarCipher("yzYZ-[]-abc", 27)
  fmt.Println(ciphered)
}

func camelcase(s string) int32 {
  count := 1
  for _, charCode := range s {
    if charCode >= 65 && charCode < 97 {
      count++
    }
  }

  return int32(count)
}

func caesarCipher(s string, k int32) string {
  ciphered := ""

  if k > 26 {
    k = k % 26
  }

  for _, charCode := range s {
    if charCode < 65 || charCode > 122 || (charCode > 90 && charCode < 97) { // letters range
      ciphered += string(charCode)
    } else {
      ciphered += cipherChar(charCode, k)
    }
  }

  return ciphered
}

func cipherChar(charCode rune, k int32) string {
  if charCode >= 97 { // start of lowercase letters
    return cipherLowChar(charCode, k)
  } else {
    return cipherUpChar(charCode, k)
  }
}

func cipherLowChar(charCode rune, k int32) string {
  newCharCode := int32(charCode) + k
  if newCharCode > 122 {
    newCharCode = 96 + (newCharCode % 122)
  }
  return string(newCharCode)
}

func cipherUpChar(charCode rune, k int32) string {
  newCharCode := int32(charCode) + k
  if newCharCode > 90 {
    newCharCode = 64 + (newCharCode % 90)
  }
  return string(newCharCode)
}
