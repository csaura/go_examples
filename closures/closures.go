package main

import "fmt"

// returns another function, which we define anonymously in the body of intSeq
// The returned function closes over the variable i to form a closure
func intSeq() func() int {
  i := 0

  return func() int {
    i += 1
    return i
  }
}

func main() {
  // We call intSeq, assigning the result (a function) to nextInt
  // This function value captures its own i value, which will be updated each time we call nextInt.
  nextInt := intSeq()

  // every time is calles the state is changed
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // new state initialization
  nextInt = intSeq()
  fmt.Println(nextInt())
}
