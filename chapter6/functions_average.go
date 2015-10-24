package main

import "fmt" 
// This function calculates averages. 
// I declare a function named "average." 
// The first parameter is named values. It is a slice of the type float64. 
// The return value is also of the type float64. 
func average(values []float64) float64 {
// A variable total is declared and initialized. 
    total := 0.0
// A variable "v" passes through "values." 
// Every time it passes through "values" it adds all the previous values. 
// _ after for shows that the program doesn't need an iterator.  
    for _, v := range values {
        total += v
    }
// This statement returns the value stored in the total divided by the length of the slice, as measured by the number of its elements. 
    return total / float64(len(values)) 
}

func main() {
// "values" takes five parameters of type float64. 
    values := []float64{1, 3, 4, 60, 89, 90}
// Here is another slice of values. 
    distance := []float64{10, 20, 100}
// A parameters are input into the function average and the result is displayed on the screen. 
    fmt.Println(average(values)) 
// This time another parameter is input. 
    fmt.Println(average(distance)) 
}
