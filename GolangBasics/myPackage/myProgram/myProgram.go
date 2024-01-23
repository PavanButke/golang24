package myProgram

import (
	"fmt"
 
)

func checkDT(){
	var a interface{}
	a="This a data itself."

	switch t: a(type) {
	case int64:
		fmt.Println("Type is an integer: ", t )
	
	case string:
		fmt.Println("Type is an string: ", t )

	case boolean:
		fmt.Println("Type is an boolean: ", t )
	
	default:
		fmt.Println("Idk , the data type!")
	}
}