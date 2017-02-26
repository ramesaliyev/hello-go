package main

import "fmt"

func main() {
    fmt.Println(biggest(10, 20, 1, 3, 5, 99))
}

func biggest(nums ...int) int {
    var most int = 0;

    for _, i := range nums {
        if i > most {
            most = i
        }
    }

    return most
}