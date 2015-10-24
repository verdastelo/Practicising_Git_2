package main

import "fmt" 

func main() {
    welcome := "This program converts Fahrenheit values in Celsius.\n " 
    prompt := "Please enter a value: "
    fmt.Println(welcome, prompt) 
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit) 
    celsius := (fahrenheit - 32) * (5.0/9.0)
    fmt.Println(celsius) 
}
