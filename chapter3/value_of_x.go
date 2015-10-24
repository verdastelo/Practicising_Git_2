package main

import "fmt" 

func main() {
    const x int = 20
    fmt.Println(x) 
    two() 
    fmt.Println(x) 
}

func two() {
    var x int = 21
    fmt.Println(x)
}
