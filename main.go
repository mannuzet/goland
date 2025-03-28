package main 

import "fmt"

func main(){
	var number1 int
	var number2 int
 	fmt.Println("Type one number:")
 	fmt.Scan(&number1)
	fmt.Println("Type another number")
	fmt.Scan(&number2)
    
 fmt.Println("the sum is:", number1 + number2)
 fmt.Println("the subtration is:", number1 - number2)
 fmt.Println("the mutiplication is:", number1 * number2)
 fmt.Println("the divison is:", number1 / number2)
 fmt.Println("the rest of the division is:", number1 % number2)
}
