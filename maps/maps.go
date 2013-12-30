// Maps are called hashes or dicts in other languages

package main

import "fmt"

func main() {
     // To create an empty map, use the builtin make:
     // make(map[key-type]val-type)
     m := make(map[string]int)

     m["k1"] = 7
     m["k2"] = 13

     fmt.Println("maps: ", m)
     
     v1 := m["k1"]
     fmt.Println("v1 :", v1)

     // len returns the number of key/value pairs when called on a map.
     fmt.Println("len: ", len(m))

     // delete removes key/value pairs from a map

     delete(m, "k2")
     fmt.Println("map: ", m)

     // second return value when getting a value from a map indicates if the key was present in the map.
     // This can be used to disambiguate between missing keys and keys with zero values like 0 or ""
     _, prs := m["k2"]
     fmt.Println("prns", prs)

     n := map[string]int{"foo": 1, "bar": 2}
     fmt.Println("map :", n)
}