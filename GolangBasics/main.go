package main

import (
	"fmt"
 "reflect"
)

const(
	Name = "Pavan"
	DOB  = "1307"
	TYPE = "Sapien"
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
	scanOps()

}	

func test(){
	fmt.Println("Testing started!")
}
func sum(num1 int , num2 int , num3 int) (int, int , int){
	res := num1 +num2 +num3 
	return res,0,0;
}

func scanOps(){
	var userInput string
	fmt.Println("Please enter Age: ")
	fmt.Scan(&userInput)
	fmt.Printf("User %s is born on %s and his age is %s of type %s",
		Name,
		DOB,
		userInput,
		TYPE,
		
	)
}