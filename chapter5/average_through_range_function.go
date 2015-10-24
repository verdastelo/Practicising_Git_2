package main

import "fmt" 

func main() {
    var x [5]float64
    x[0] = 1989
    x[1] = 1979
    x[2] = 1999
    x[3] = 1969
    x[4] = 1959
    
    var total float64 = 0

    for _, value := range x {
        total += value
    }

    fmt.Println(total/float64(len(x))) 
    }


