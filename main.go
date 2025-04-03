package main

import "fmt"

func main() {
    var ages = [4] int{17, 16, 20, 40}
    names := [4] string{"Manu", "Belly", "Isa", "Ju"}

    fmt.Println(ages)
    fmt.Println(names)

   names[0] = "Manuzett" 
   fmt.Println(names)

   // Slice
   var score = []int{100, 200, 300, 400}
   fmt.Println(score)

   score [1] = 2
   fmt.Println(score, len(score), cap(score))

   rangeOne := score[1:3]
   fmt.Println(rangeOne)

   rangeTwo := score[2:]
   fmt.Println(rangeTwo)

   rangeThree := score[:3]
   fmt.Println(rangeThree)

   var superheros = []string{"Deadpool", "Spider-man", "Batman"}
   fmt.Println(superheros)

   superheros = append(superheros, "Ironman", "Thor")
   fmt.Println(superheros, len(superheros), cap(superheros))
 

}