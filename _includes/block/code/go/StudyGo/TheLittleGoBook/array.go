package main

import "fmt"

// Arrays are efficient but rigid.

// In Go, like many other languages, arrays are fixed.
// Declaring an array requires that we specify the size,
// and once the size is specified, it cannot grow.
func main() {
    var a [10]int
    a[0] = 10
    // we can use range to iterate over the array.
    for _, v := range a {
        fmt.Printf("%d, ", v)
    }
    fmt.Println()
    // We can use len to get the length of the array.
    fmt.Printf("The length of array a is: %d.\n", len(a))

    scores := [4]int{ 100, 99, 98, 97 }
    for _, score := range scores {
        fmt.Printf("%d, ", score)
    }
    fmt.Println()
    fmt.Printf("The length of array scores is: %d.\n", len(scores))
    fmt.Println("done.")
}

// output:
//
// 10, 0, 0, 0, 0, 0, 0, 0, 0, 0,
// The length of array a is: 10.
// 100, 99, 98, 97,
// The length of array scores is: 4.
// done.
