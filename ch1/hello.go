package main //specifies that this main package
//starting of the go program

import "fmt" //importing format package

//package level declaration
//we can have package level delcarations
const f string ="This is a const ";

//start of the program
func main() { //function declaration 
	// single line comment
	/* multi line comment  */

	//Variable declartion
	// var x,y int ;
	//function scope
	//Declaraiton List
	var (
		x int ;
		y='j' ;
		z float32;
	);


	a,b:=2,3; //opt out var interference
	c,x:=4,3;

	//this prints in new line 

	x=int(z);
	fmt.Println(x, y,z,a,b,c,f);//funciton calling
}

/*
Go has lot of ways to declare variables

Verbose way (explict type)
var x int =  10

Still same (implict type)
var x = 10
var x,y int = 10,20

			or
Skip assignment (by default value will be zero)
var x int
var x,y int


if we skip assignment then give explicitily type

Multile assignment
var x,y = 10,"hello";

Declaration List
	var (
		x int ;
		y='j' ;
		z float32;
	)

Go also supports short declaraiton and assignment format
within a function only 
:=

if we want to declare that at global level 
we need to use var

	a,b:=2,3;

we can also re assign existing variables too
with one condition that 
atleast one new variable should be use 


best to use within fucntion := and out side var



avoid declariing variables at package level


var and const declartions are same

runtime const declarations are invalid in go
var x int =10;
const x int = x*5; //illegal
we can specify the type and untype the type for go constants



go is call by value

there should not be any un used variables 


*/
