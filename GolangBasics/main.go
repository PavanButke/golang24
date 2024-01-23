package main

import (
	"fmt"
 //"reflect"
)

const(
	Name = "Pavan"
	DOB  = "1307"
	TYPE = "Sapien"
)
func main()  {	

	var a=12 
	var ptr1 *int

	var b=21
	var ptr2 *int

	ptr1 = &a
	ptr2 = &b 

	fmt.Printf("This is the value %v and %v", a , b)
	fmt.Printf("\nThis is the  adddress of %v and %v", &a , &b)
	fmt.Printf("\nThis is the  adddress of %v and %v", ptr1 , ptr2)
	fmt.Printf("\nThis is the  values of  a and by ptr1 & ptr2 %v and %v", *ptr1 , *ptr2)


	c:= *ptr1 + *ptr2
	fmt.Println("\nThe sum of values  ",c)

	// var a interface{}
	// a="This a data itself."

	// switch t:= a.(type) {
	// case int64:
	// 	fmt.Println("Type is an integer: ", t )
	
	// case string:
	// 	fmt.Println("Type is an string: ", t )

	// case bool:
	// 	fmt.Println("Type is an boolean: ", t )
	
	// default:
	// 	fmt.Println("Idk , the data type!")
	// }
	// fmt.Println("Hello Computers")
	// var x int
	// var y =10
	// z:=10

	// fmt.Printf("res: x is typeof %v and y is type of %v and z is type of %v\n",
	// reflect.TypeOf(x),
	// reflect.TypeOf(y),
	// reflect.TypeOf(z))
	// test()

	// fmt.Println(sum(x,y,z))
	// scanOps()

}	

// func test(){
// 	fmt.Println("Testing started!")
// }
// func sum(num1 int , num2 int , num3 int) (int, int , int){
// 	res := num1 +num2 +num3 
// 	return res,0,0;
// }

// func scanOps(){
// 	var userInput string
// 	fmt.Println("Please enter Age: ")
// 	fmt.Scan(&userInput)
// 	fmt.Printf("User %s is born on %s and his age is %s of type %s",
// 		Name,
// 		DOB,
// 		userInput,
// 		TYPE,
		
// 	)
// }