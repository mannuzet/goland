package main

import "fmt"

func main() {
    var user string
    var password string
   
    fmt.Print("Enter your username")
    fmt.Scanln(&user)

    fmt.Print("Enter your password")
    fmt.Scanln(&password)


    if user == "mannu" && password == "123017"{
    fmt.Println("Allowed access")

    }else{ 
        fmt.Println("Access denied")
    
   

    
 }
    
    
}

