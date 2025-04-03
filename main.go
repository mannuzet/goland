package main

import "fmt"

func main() {
    var ages = [4] int{17, 16, 20, 40}
    names := [4] string{"Manu", "Belly", "Isa", "Ju"}
    fmt.Println(ages)
    fmt.Println(names)
   names[0] = "Manuzett" 
   fmt.Println(names)
 }

