package greetings

import "fmt"

//Hello is the func name; 
//seems it goes [name of the variable][type of the variable]
//return type goes at the end
//In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.
func Hello(name string) string {
	//variable declaration can be done two ways
	/*first way
	var message string
	message = fmt.Sprint("Hi... %v", name)
	second way with := */

	message:= fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
