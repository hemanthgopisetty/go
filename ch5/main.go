package main

import "fmt"

type Person struct {
	firstName string
	lastName  *string
}

func makePointer[T any](t T) *T {
	return &t
}
func main() {

	var y = Person{
		firstName: "Hemanth",
		lastName:  makePointer("Gopisetty"),
	};
	var x *Person=&y // this is pointer to strutct
	fmt.Println(x,&x,&y)
	x.firstName="Raju"
	fmt.Println(*y.lastName)
}

/*
Comparing Classes in other languages with the pointesr in Go

Pointers in go makes the go code faster

Pointer is a variable that stores the location of another variable where it is stored

var x int32 = 10  // this is 32 bit integer (storing 32 bits need 4bytes )(1 byte = 8 bits)
var y bool = true // this is 1 bit size of variable (but the smallest memory location is 1 byte)
// this boolean will take 1 byte to represent the false or true(actually u need 1 bit right to represent)

location of x is 1234(4 byte)//starting at address 1 and ending at 4
location of y is 5(1 byte)//starting and ending at address 5

each variable is stored in one or more contigous memory location called addresses

diff type of variable diff amount of memmory

but not what pointer it;s it will take same amount of memory


A pointer is a variable that stores location of another variable in memory
 A pointer holds a number , that is nothing but address of another memory

no value of pointer is nil
maps, and slices are implemented using pointers

nil is an untyped identifier that represents
the lack of value for certain types

"go has garbage collector where as c and c++ langauges doesn't have
and go has borrowed it's pointer syntax from c and c++"

"no pointer airthmetic in Go unlike in C and C++"


&variable give the location of variable int the memory
we can access the data using indirection operator * on the pointer to access the location data

this process is known as dereferencing

before dereferecing a pointer pls make sure it;s not nil

var ponterTox* int -> pointer type decalration
x:=&y -> auto assignment type

new -> is operator that creates a memory of the type and return's a pointer

var x = new(int) -> return's a pointer

we can also store the struct address and dereference with indirection operator

to convert a constant literal to a pointer use helper generic function which returns the


Pointers are actually the familiar behaviour of classes

Go lang is pass by value not at all pass by reference 
by default unlike other languages like if we take java or python if we pass it to another funnctin
we can see modified change of the varaible if we changed

go gives comfortness compared to others
only by passing pointers u will have this effect which is pass by reference but still is Go lang is call
by value , when a pointer parameter is passed to a function
the function copy the variable of the pointer and the value of the pointer

pointers indicate mutable parameters

Immutability is the best way to construct the software

we should not pass nil pointer to a function parameter

suppose if we pass and assign a new pointer address to that
then it won;t modify the new variable 

copying a pointer has two implications
1)Should not pass or copy the nil pointer
2)if we want to change the value then derefence and set the value of passed pointer

Pointers can create confusion among variables
and hard to create the data flow 
and loads the garbage collector

and don;t pass heavy pointer since go is copy of value 
it's very hard for garbage collector to clean the memory

The only time we should use pointers parameters it modify a variable when the function expects
an interface

Pointers passing performance
suppose if struct is very large then passing pointers by struct 
is more performance since go is call by value

Behaviour of returning a pointer and returning a value is performance depends

if the data is very huge to pass a function then use pointers this will really impact the performance

Don;t pass maps becuase maps are impemented with the pointers
and we are passing pointer means there is a scope for mutability


so use structs for the performance and security perspective

Slice And Pointer :- 
If we pass slice as pointer we can't change size of the pointer
whereas we can change the copied slice but it  won;t reflect
in the orginal slice

this is happening because dude go is call by value

but we can change the content's of the orginal slice till the length
of the slice


*/
