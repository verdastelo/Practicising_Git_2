package main

import "fmt" 

func add (numbers ...int) int {
    total := 0
    for _, i := range numbers {
        total += i
    }
    return total
}

func main() {
    data := []int{1, 2, 4, 5, 6}
    fmt.Println(add(1, 3, 5, 7, 9)) 
// Notice the syntax. 
    fmt.Println(add(data...)) 
} 
