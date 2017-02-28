package main

import "fmt"

func main() {
    // multi-dimensional array
    var multiarr [3][3][3]int
    multiarr[0][0][0] = 1
    fmt.Println(multiarr)

    // example of not being reference type
    var intarray [10]int

    fmt.Println(intarray)

    intarray[1] = 1
    intarray[2] = 2

    fmt.Println(intarray)

    // array is not reference type so its gonna PASS BY VALUE
    change(intarray)

    // dont changed
    fmt.Println(intarray)

}

func change(inta [10]int) {
    inta[3] = 3
    inta[1] = 10

    // different that original
    fmt.Println(inta)
}