package main

import (
	"fmt"
 "reflect"
)

func main()  {
	fmt.Println("Hello Computers")
	var x int
	var y =10
	z:=10

	fmt.Printf("res: x is typeof %v and y is type of %v and z is type of %v\n",
	reflect.TypeOf(x),
	reflect.TypeOf(y),
	reflect.TypeOf(z))
	test()
	fmt.Println(sum(x,y,z))

}	

func test(){
	fmt.Println("Testing started!")
}
func sum(num1 int , num2 int , num3 int) (int, int , int){
	res := num1 +num2 +num3 
	return res,0,0;
}