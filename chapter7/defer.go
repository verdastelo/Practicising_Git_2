package main 

import "fmt" 

func first() {
    fmt.Println("Hodiaŭ mi komencis legi libron Esperante.") 
}

func second() {
    fmt.Println("Книга Карлсона хорошая.") 
} 

func main() {
    defer second() 
    first() 
} 
