package main

import "fmt"

func main(){

var x = max(1,2)

fmt.Println(x)

}


func max(num1,num2 int)int{

var result int

if(num1 > num2){

result = num1
}else{
result = num2
}

return result

}

