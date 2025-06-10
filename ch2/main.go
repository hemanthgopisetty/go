package main
import "fmt"
func main(){
 
	var x =[2][3]int{{1,3},{}};
	x[1][0]=8
	fmt.Println(len(x));

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
 
*/