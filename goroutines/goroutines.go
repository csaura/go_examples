package main

import "fmt"

// A goroutine is a lightweight thread of execution.
func f(from string) {
  for i := 0; i < 3; i++ {
    fmt.Println(from, ":", i)
  }
}

func main() {
  //  running it synchronously.
  f("direct")

  //new goroutine will execute concurrently with the calling one
  go f("goroutine")

  // start a goroutine for an anonymous function call.
  go func(msg string) {
    fmt.Println(msg)
  }("going")

  var input string

  //Our two goroutines are running asynchronously in separate goroutines now, so execution falls through to here. This Scanln code requires we press a key before the program exits.
  fmt.Scanln(&input)
  fmt.Println(input)
}

