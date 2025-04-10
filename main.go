package main

import (
    "fmt"
)

func main() {
    ages := 45
    fmt.Println(ages <= 50)
    fmt.Println(ages >= 50)
    fmt.Println(ages == 50)
    fmt.Println(ages != 50)

    if ages < 50 {
      fmt.Println("Less than 30 years")
    } else if ages < 40 {
      fmt.Println("Less than 40 years")
   } else { 
      fmt.Println("Isn't less than 40 years")
   }
   names := []string{"Emanuela", "Isabelly", "Isadora"}
   for index, value := range names {
      if index == 1 {
         fmt.Println("Continue after the position", index, "and value", value)
         continue
      }
      if index > 2 {
         fmt.Println("Leave after" ,index)
         break
      }
      fmt.Println("Value:" ,value)
   }
}