package main

import "fmt" 

func main() { 
    elements := make(map[string]string)
    elements["H"] = "Hydrogen" 
    elements["He"] = "Helium" 
    elements["Li"] = "Lithium" 
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron" 
    elements["C"] = "Carbon" 
    elements["N"] = "Nitrogen" 
    elements["O"] = "Oxygen" 
    elements["F"] = "Fluorine" 
    elements["Ne"] = "Neon"


    fmt.Println(elements["Li"]) 

// Here "name" is assigned the value stored under the key "O."
// If the value does not exist, "name" is the value type's zero value. (An empty string in the present case.) 
// The second value "ok" is a bool.
// "ok" is true if the key exists in the map, and false if it does not.      
// Assign the value under the key "O" to name and print it if the key exists. 
    if name, ok := elements["O"]; ok {
        fmt.Println(name, ok) // This "ok" is optional. 
    }
} 
