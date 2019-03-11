package main

import (
    "fmt"
    "os"

    "github.com/NikolaBergo/fizzbuzz"
)


func main() {
    fmt.Printf("Welcome to fizzbuzz!\nPlease, write your data to 'inputdata' file and press ENTER\n")

    var any byte
    fmt.Scanf("%c", &any)

    r, _ := os.Open("inputdata")
    w, _ := os.Create("outputdata")

    err := fizzbuzz.FizzBuzz(r, w)

    if err != nil {
        fmt.Printf("Ooop-s, something went wrong, check your input, error: ")
        fmt.Print(err)
        fmt.Print("\n")
    }

    r.Close()
    w.Close()
}
