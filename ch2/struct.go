package main

import "fmt"

func practiceStruct() {
	type person struct{
		name string
		age int 
		pet string

	};

	hemanth := person{
		name: "hemant",
		age: 9,
		pet: "dog",
		
	};
	fmt.Println(hemanth)
}

/***********************************************************
all values in the map must be of the same type
for this reason maps are not suitable for all operations
maps are not ideal to pass data from function to function

all values of map mut be the same type

structs are suitable for gouping all data together
go doesn;t have classes
becuase it doesn't support inheritance

instead it have structs

	type person struct{
		name string
		age int 
		pet string
		
	}
	struct type is defined by this type keyword 
	and struct keyword , {}
    between the types no commas are needed

	you can define struct type inside or outside
	of a function

	struct defined in function 
	can be accesible in functions only
    has scopes

	once struct type is defined
	then 
	var hemanth person ; var declaration
	hemanth:=person{}

    	hemanth := person{
		name: "hemant",
		age: 9,
		pet: "dog",
		
	};

	we can specify the fields in any order
	accesed with a dot notation
	write with a dot notation


	Anonymous Structs

	var implements a struct type
	without first giving the struct type name

	var person struct{
	   person string 
	   age int
	   name string
	}
    same goes for string literal

	these are anonymous structs


	Comparable Structs
	wether a struct is compare based on the type of

	structs are composed of comparable type


	go deosn't allow comparision of different  variables of primitive types
    same goes foro structs to

	strutsc that are entieryl composed of comparable
	types are comparable


	go allows converion only when order , names , type are some
	envent though u can;t compare

	



***********************************************************/
