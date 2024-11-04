//main package groups most commonly used functions I guess
package main

//importing a package from the main package also must be done I guess
//this package contains functions for formatting text, including printing to the console; a standard library package
import (
	"fmt"

	"gopherit/greetings"
)

import "rsc.io/quote"

//this is how you declare the main function
func main() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())

	message:=greetings.Hello("Mils")
	fmt.Println(message)
}

/*additional notes:
to run you write 'go run .'
check what else you can do with 'go help'

there are different packages that can be installed from pkg.go.dev*/
