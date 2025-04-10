package main

import (
"fmt"  
"strings"
"sort"
)

func main() {
	greeting := "Hello my friends"
    fmt.Println(strings.Contains(greeting, "friends"))
    fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
    fmt.Println(strings.ToUpper(greeting))
    fmt.Println(strings.ToLower(greeting))
    fmt.Println(strings.Index(greeting, "my"))
    fmt.Println(strings.Split(greeting, "  "))
    ages := []int {50, 80, 10}
    sort.Ints(ages)
    fmt.Println(ages)
    index := sort.SearchInts(ages, 50)
    fmt.Println(index)
    names := []string {"Sabrina", "Olivia", "Lana"}
    sort.Strings(names)
    fmt.Println(names)
    fmt.Println(sort.SearchStrings(names, "Sabrina"))
   x :=0

   for x < 5 {
   fmt.Println(x)
   x++
}

for i:=0; i < 5; i++{
    fmt.Println("for 2:", i)
}

for i:=0; i <len(names); i++{
    fmt.Println(names[i])
}
for index, value := range names{
    fmt.Println("the index is:", index, "and the value is:", value)
}

for index, value := range ages{
    fmt.Println("the index is:", index, "and the value is:", value)
}

superheros := []string{"Deadpool", "Spider-man", "Batman"}

for index, value := range superheros{
    fmt.Println("The hero number is:", index, "the hero name is:", value)
}
}