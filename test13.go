package main

import "fmt"

func main(){

var greeting = "Hello World";
fmt.Printf("normal string:")
fmt.Printf("%s", greeting)
fmt.Printf("\n")
fmt.Printf("hex bytes:")

for i:= 0 ; i < len(greeting); i++{
fmt.Printf("%x",greeting[i])
}


}

