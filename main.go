package main 

import "fmt"

func main(){
	var number int
	fmt.Print("Type a number:")
	fmt.Scan(&number)
	fmt.Println("The typed number is", number)
}