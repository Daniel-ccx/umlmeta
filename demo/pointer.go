package main

import (
    "fmt"
)

func main() {
    it := 1
    fmt.Println("initial:", it)
    var ip1 *int
    ip1 = &it
    fmt.Println("zeroPtr:", it)
    fmt.Println("zeroPtr:", ip1)

    fmt.Println("pointer:", &it)
}
