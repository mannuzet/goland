package main

import (
    "fmt"
)

func sayGreeting(name string){
fmt.Println("Hii", name)
}

func addNumber(number1 int, number2 int) int{
    return number1 + number2
}
func main() {
sayGreeting("Sabrina")
result := addNumber(10, 20)
fmt.Println(result)

}