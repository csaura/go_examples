package main

import "fmt"

// takes two ints and returns their sum as an int
func plus(a int, b int) int {
  // requires explicit returns
  return a + b
}

func main() {
  res := plus(1, 2)
  fmt.Println("1+2 =", res)
}

