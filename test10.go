package main
import "fmt"

func main(){
var x,y string = "mahesh","sumesh"

a, b := swap(x,y)
fmt.Println(a,b)

}


func swap(a,b string)(string,string){
return b,a
}
