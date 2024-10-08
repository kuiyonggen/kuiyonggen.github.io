package main

import "fmt"


func main() {
    m := make(map[string]int)
    m["alice"] = 10

    // Get the value if exists, otherwise return 
    // the default value according to the value type 
    // and false.
    v, exists := m["jim"]
    fmt.Printf("v = %d, exists = %v.\n", v, exists)

    total := len(m)
    fmt.Printf("total: %d.\n", total)

    delete(m, "alice")
    total = len(m)
    fmt.Printf("total after deleting: %d.\n", total)

    // Declare a map as a composite literal.
    m1 := map[string]int{
        "alice": 10,
        "alvin": 6,
    }

    // Iteration over maps isn't ordered.
    // Each iteration over a lookup will 
    // return the key value pair in a random order.
    for k, v := range m1 {
        fmt.Printf("k: %s, v: %d.\n", k, v)
    }
    
    fmt.Println("done.")
}
