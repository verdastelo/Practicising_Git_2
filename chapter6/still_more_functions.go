package main

import "fmt" 
// Tiu ĉi funkcio vokas la funkcion f1(). 
func main() {
    fmt.Println(f1()) 
}
// When f1() is called, it returns f2(), which incidentally is another function. 
func f1() int {
    return f2() 
}
// Calling f2() returns 1. 
func f2() int {
    return 1
}
