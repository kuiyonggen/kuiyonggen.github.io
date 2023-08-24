package main

import "fmt"

// A slice is a lightweight structure 
// that wraps and represents a portion of an array.
func printSlice(name string, a []int) {
    fmt.Printf("%s: ", name)
    for _, v := range a {
        fmt.Printf("%d, ", v)
    }
    fmt.Println()
    fmt.Printf("%s: %v.\n", name, a)
    fmt.Printf("The capacity of %s is: %d.\n", name, cap(a))
    fmt.Printf("The length of %s is: %d.\n", name, len(a))
    fmt.Println()
}

func main() {
    fmt.Println("Create a slice by declaration.")
    var s []int
    printSlice("s", s)

    fmt.Println(`Create a slice by declaration and assignment 
without a length comparing with array declaration.`)
    s0 := []int{100, 99, 98, 97, 96}
    printSlice("s0", s0)

    // The capacity is the length from the start to the end of original array.
    s01 := s0[2:3]
    printSlice("s01", s01)

    s02 := s0[2:4]
    printSlice("s02", s02)

    s03 := s0[1:4]
    printSlice("s03", s03)

    a := [10]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
    sa1 := a[2:4]
    printSlice("sa1", sa1)
    sa2 := sa1[1:2]
    printSlice("sa2", sa2)

    // copy() doesn't increase the capacity of slice.
    sb := make([]int, 4)
    copy(sb, a[:3])
    printSlice("sb after copying a[:3]", sb)
    copy(sb, a[:4])
    printSlice("sb after copying a[:4]", sb)
    copy(sb, a[:])
    printSlice("sb after copying a[:]", sb)

    fmt.Println("Create a slice by make([]int, 4):")
    // Initialize a slice with a length of 4
    // and a capacity of 4.
    // The length is the size of the slice,
    // while the capacity is the size of the underlying array.
    s1 := make([]int, 4)
    s1[3] = 999
    printSlice("s1", s1)

    s1 = append(s1, 100)
    printSlice("s1 after appending 100", s1)

    fmt.Println("Create a slice by make([]int, 0, 4):")
    // Create a slice with a length of 0
    // but with a capacity of 4.
    s2 := make([]int, 0, 4)
    printSlice("s2", s2)

    // One way to expand a slice is via "append"
    s2 = append(s2, 100)
    printSlice("s2 after appending 100", s2)

    // Append four more values, causing the capacity 
    // to increase the number by its current size.
    s2 = append(s2, 99, 98, 97, 96)
    printSlice("s2 after appending four more values", s2)

    // Append five more values
    s2 = append(s2, 95, 94, 93, 92, 91)
    printSlice("s2 after appending five more values", s2)

    fmt.Println("done.")
}
