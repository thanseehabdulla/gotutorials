package main

import "fmt"

/* global variable declaration */

var a int = 20

func main(){

var a int = 10
var b int = 20
var c int = 0

fmt.printf("value of a in main() is %d\n", a )
c = sum( a, b)
fmt.printf("value of c in main() = %d", c)

}

func sum(a, b int) int {
fmt.printf("the value of a in sum() = %d\n", a)
fmt.printf("value of b ")

return a + b;

}
