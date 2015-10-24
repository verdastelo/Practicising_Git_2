package main

import "fmt" 

func main() {
    list := "List of numbers divisible by 3 and less than 100." 
    fmt.Println(list) 
    for i := 1; i <= 100; i++ {
        if ((i % 3) == 0) && ((i % 5) == 0) {
            fmt.Println("Fizzbuzz") 
        } else if i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz") 
        } else {
            fmt.Println(i) 
        }
    }
}
