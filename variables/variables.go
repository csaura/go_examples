package main

import "fmt"

func main() {
  var a string = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  // Variables declared without a corresponding initialization are zero-valued
  var e int
  fmt.Println(e)

  // The := syntax is shorthand for declaring and initializing a variable,
  f := "sort"
  fmt.Println(f)

}
