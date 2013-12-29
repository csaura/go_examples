package main

import "fmt"

func main() {
     var a [5]int

     // By default an array is zero-valued, which for ints means 0s
     fmt.Println("emp:", a)

     a[4] = 100
     fmt.Println("set:", a)
     fmt.Println("get:", a[4])

     // len() returns the length of an array
     fmt.Println("len:", len(a))

     // syntax to declare and initialize an array in one line
     b := [5]int{1, 2, 3, 4, 5}
     fmt.Println("dcl:", b)

     var twoD [2][3]int

     for i := 0; i < 2; i++ {
     	 for j := 0; j < 3; j++ {
	     twoD[i][j] = i + j
	 }
     }

     fmt.Println("2d: ", twoD)
}