package main

import (
    "fmt" 
   // "sort" 
) 

func greatest_number(num ...int) int {
    fmt.Println(num) 
    s := []int 
    append(s, num) 
    fmt.Println(s) 
    return 0
}

func main() {
    greatest_number(1, 2, 3) 
} 
