package main

import "fmt"
import "time"

//We can use channels to synchronize execution across goroutines. Here’s an example of using a blocking receive to wait for a goroutine to finish.

// The done channel will be used to notify another goroutine that this function’s work is done.
func worker(done chan bool) {
  fmt.Print("working...")
  time.Sleep(time.Second)
  fmt.Print("done")

  //Send a value to notify that we’re done.
  done <- true
}

func main() {
  //Start a worker goroutine, giving it the channel to notify on
  done := make(chan bool, 1)
  
  go worker(done)

  //Block until we receive a notification from the worker on the channel.
  a := <- done

  fmt.Println(a)
}
