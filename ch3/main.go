package main

import "fmt"
func shadow_varaibles() {
	x := 10 // function block
	y := 20
	if x > 5 {
		fmt.Println(x, y)
		y, x := 5, 15 // contained block
		fmt.Println(x, y)
	}
	fmt.Println(x, y)

}
func main() {
	// shadow_varaibles()
	// if n := 5; n == 0 {
	// 	fmt.Println("If block")
	// 	fmt.Println(n)
	// } else {
	// 	fmt.Println("else block")
	// 	fmt.Println(n)
	// }
	// forLoop();
	// switchS();
	gotoS();
}

/*
We can declare variables in go in lots of place

each where declaration is called a block

variables , constants , types , functions declared
outside of any functions are placed in the package block

printing and some math functions importing
from other file are in file scope


all varaiables defined in the top level
of the function
are in FUnction level


within a function {}
every set of braces defines another
block


control strutures define another
block of their own set

you can access any variable outer block from
within inner block


shadowing variable is variable
that has the same name as a variable in the containing
block

as long as the shadowed variable exists
you cannot access the shadowed variable

:= this declaration when used for multi variable
assignment

there left side should be one old variable followed by
new variaable but as we can see in shadow if statement

it will use only the varaibles when containing block has
shadowing packagin names

there are predeclared identifiers

they have universal block
we can also redefine the identifiers

we must be very careful when
we are redeclaring the idenfiters

go vet does not report that as a error



if condition{
 ...statements
}
 go add's parantheese aroud the condtion


go adds one more ability to
if and else other than other languages
if u declare one variable
that u can use in else also
to make availabiltiy for else
we need to create that in out if and else
in other languages
but here if we check we can see that
	if n := 5; n == 0 {
		fmt.Println("If block")
		fmt.Println(n)
	} else {
		fmt.Println("else block")
		fmt.Println(n)
	}

	after if and else there is no scope
	for n 
	and n is undefined

	before the compartion of
	if statement
	we can put any simple statement

	variables declared as part of if 
	and else
	will shadow the varaibles too

	
*/
