package main

import (
    "fmt"
)

func toDivide(dividend int, divider int) (int, string){
if divider == 0 {
     return 0, "Error in division by 0"
}
return dividend / divider, "No error"
}

func basicOperation(a int, b int)(int, int, int){
    sum := a + b
    multiplication := a * b
    subtraction := a - b

    return sum, multiplication, subtraction
}

func main() {
    result, erro := toDivide(10,2)
    
    if erro != "No error"{
        fmt.Println(erro)
    } else {
        fmt.Println("The result is:", result,)
    }
   
    sum, mult, sub := basicOperation(10,2)
    fmt.Println(sum)
    fmt.Println(mult)
    fmt.Println(sub)

}