package main
import "fmt"
type cat struct{
	name string
	weight int 
}
func fillCat(myCat cat){
	fmt.Println(myCat)
}

func add(a int , b int)int {
	return a+b;
}

func multiply(a...int)int{
	x:=1
	for _,v:=range(a){
		x*=v
	}
	return x
}
func returnMultipleByNames()(a,b,c int){
	fmt.Println(a,b,c)
	return 0,0,0
}
func returnMultiple()(int,int,string){
	return 1,2,"Hemanth"
}
func main(){
	// fillCat(cat{
	// 	name : "white",
	// })

	// fmt.Println(multiply(1,2,3,4,5,6,7,8,9));
	// fmt.Println(multiply(1,2));
	// a:=[]int{3,5}
	// fmt.Println(multiply(a...));
	// fmt.Println(multiply([]int{1,2,3,4,5}...));

	// fmt.Println(returnMultiple())
	// fmt.Println(returnMultipleByNames())

	function()
}
/*
every go program starts from main function

func add(a int , b int)int {
	return a+b;
}

this is how a function looks like in go

if u have same types in parameter we can declare like this


Using struct to simulate to name and optional parameters

type cat struct{
	name string
	weight int 
}
func fillCat(myCat cat){
	
}


variadic parameters and slices

if we pass as many as variables to a function of same type 
then we should use varaidic parameter

and the postion of this parameter in this function definition 
is last

that variadic parameter is noting but slice of that variadic type

func multiply(a...int)int{
	x:=1
	for _,v:=range(a){
		x*=v
	}
	return x
}

	fmt.Println(multiply(1,2,3,4,5,6,7,8,9));
	fmt.Println(multiply(1,2));
	a:=[]int{3,5}
	fmt.Println(multiply(a...));
	fmt.Println(multiply([]int{1,2,3,4,5}...));


multiple return values
from a function in go like we assign mutliple variables in single
assignment 
we can do the same from a function returning 
func returnMultiple()(int,int,string){
	return 1,2,"Hemanth"
}
	fmt.Println(returnMultiple())

	multiple return values are multiple values
	there is no destruct as in js or pyth

	we can ignore the values by naming the varaible
	by _

in addition to return more than one value from a
function

we can specify names for your return values

if supply names to return values then we are pre declaring
the variable we can use them in the function 

named return values are intialized to zero values when created
we can shadow these variables too


func returnMultipleByNames()(a,b,c int){
	fmt.Println(a,b,c)
	return 0,0,0
}

we can also return blank values but calling function
recives intitliaze values
*/