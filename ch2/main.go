package main

import "fmt"

func main() {

	//Slice
	z := make([]string, 1)
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))
	z = append(z, "Hello")
	fmt.Println(z[1])
	fmt.Println(cap(z))
	fmt.Println(len(z))
	clear(z)
	z[0]="Raju"
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))
}

/*
 Arrays Group of same element
 Arrays start with 0
 var x[3]int ; //all elements are zero
 var x[4]int =[4]int{1,2,3}; //valid size should be same
 var x int   =[4]int{4,5,6};
 var x       =[4]int{1,3:3}; //Sparse Array's

 if you don;t know how many elements
 then
 var x =[...]int{1,900:1};


 we can compare two array's using == and !==
 	var x=[...]int{1,900:3};
	var y=[...]int{1,900:3};
	fmt.Println(x==y);

 Access element like this
 x[0]
 x[1][0]=8
 Multi Dimensional Array
 var x=[2][3]int{{1,3},{}};

Lenght of the array
len(x)
	var x =[2][3]int{{1,3},{}};
	x[1][0]=8
	fmt.Println(len(x));
	 prints 2 because size with type in array considered as a type
 Arrays in go are rarely used due to it's limitation behaviour

 go considers array size of the array to be part of type of the array

 and these types must be resolved at compile time

 type conversion also won;t work here

 Arrays existed because to support one feature Slice


 Slice  to hold sequence of vlaues
 we can grow slices as needed
 length of the slice is not part of it's type

 var x=[]int{1,2,4}
 var x = [][]int{{1},{2}} -> multi dimensional array
 var x=[]int{6:100}//sparse arrays
 var x[]int ; //this is also valid

 like array's you can compare slices since there are memory layout are going to be different
 but you can compare that with nill
 nill is like null but different compared to other languages

 since in go 1.21 slices package in standard libaray includes functions for compare

 	//Slice
	x:=[]int{1,2,3,4,5,6};
	y:=[][]int{{1,2},{3,4}};
	fmt.Println(x==y);
	invalid operation

	we can make compare this slice with nill only

	Each SLice will have Capacity to store new elmenents


	like len , we have cap , and append copies new values and returns new array each tiime we do append

	this makes go very slow

	go runtime allocate new memory and copy and return the new memory

	since go is pass by value we have to do like this dudee

	two ways of declaring slice

	var x = []int{}
	var x []int (default to nill)

	we also have third way too that is make function

	nil is an identifier reprsents lack of values


	emptying the slice with clear function

	it clears the slice and the lenght will remain same dudee
    

	z := make([]string, 1)
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))
	z = append(z, "Hello")
	fmt.Println(z[1])
	fmt.Println(cap(z))
	fmt.Println(len(z))

    Empty the slice
   	//Slice
	z := make([]string, 1)
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))
	z = append(z, "Hello")
	fmt.Println(z[1])
	fmt.Println(cap(z))
	fmt.Println(len(z))
	clear(z)
	z[0]="Raju"
	fmt.Println(z)
	fmt.Println(cap(z))
	fmt.Println(len(z))


	there are many approaches to declare a slice
	prefer to use make and append 

	
*/
