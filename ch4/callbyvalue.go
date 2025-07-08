package main;
import "fmt"
func a(x int){
	x=3
}
func callbyvalue(){
	x:=2
	a(2)
	fmt.Println(x)
}

/*
go always makes copy of value
in function paraameters or assign
map and slices are pass by reference to the memory


*/