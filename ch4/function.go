package main

import "fmt"

func function() {
	x:=func(){
		fmt.Println("X")
	}
	x()
	{var x func() = func(){
		fmt.Println("Xx")
	}
	x()}
}

/*

functions are values in go

func is a keyword
func is a type

type of a function is
	built out of the
		func keyword
		type of paramters
		return values
whole this combination is called signature


any function that has same signature can be assigned

var func printnumber() //declaration
  var printOne func();
  printOne=one;
  printOne()
func one(){
	fmt.Println(1)
}

 x:=one
 x()


 we can use them in slice,map,struct as values

 we can use type keyword to define func like struct also

	x := num
	x(1)
	type two func(x int)
	var y two = num
	y(2)

	func num(x int) {
	fmt.Println(x)
}

function types are bridge to interfaces

anonymous functions
we can define these functions inline and call the instantly
anonymous functions don't need any name they are assignable instantly
we can pass parameter as well
*/
