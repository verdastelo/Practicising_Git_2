package main

import "fmt" 

func main() {
    welcome := "This program converts feet into metres.\n" 
    prompt := "Please enter the lenght in metres." 
    fmt.Println(welcome, prompt) 
    var feet float64
    fmt.Scanf("%f", &feet) 
    metre := feet * 0.3048
    fmt.Println(metre)
}
