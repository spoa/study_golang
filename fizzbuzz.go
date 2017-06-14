package main

import (
  "fmt"
  "strconv"
)

func fizzbuzz(count int) {
  for i := 0; i < count; i++ {
    count := i + 1
    if count % 3 == 0 {
      fmt.Printf("Fizz")
    }
    if count % 5 == 0 {
      fmt.Printf("Buzz")
    }
    if count % 3 > 0 && count % 5 > 0 {
      fmt.Printf(strconv.Itoa(count))
    }
    fmt.Printf("\n")
  }
}

func main() {
  fizzbuzz(100)
}
