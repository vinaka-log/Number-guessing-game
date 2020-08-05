package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    answer := rand.Intn(10) + 1
    var guess int

    fmt.Print("Your guess? ")
    fmt.Scanf("%v", &guess)

    if answer == guess {
        fmt.Println("Bingo!")
    } else if answer > guess {
        fmt.Println("The answer is higher!")
    } else {
        fmt.Println("The anwer is lower!")
    }

    fmt.Printf("The answer is %v\n", answer)

}