package main;
import "fmt"
func deferr(){
	fmt.Println("Printing the thing")
	return 
	f:=func(){
		fmt.Println("this is a function varaible ")
	}
	defer func(){
		fmt.Println("After the exectiuon cleanup")
		f()
	}()
	fmt.Println("Printing the thing")
}
/*

defer is used to cleanup the resources in the function 
it used for cleanup process 

programs often create temporary process , resources , files or network connections
that needs to be cleaned up

if a function exits early then all these cleanup should happen
no matter wether function executed suceessfully or not

defer cleanup code // this is the syntax


Defer is used to ensure that a function call is performed later in a program’s execution

usualy for the purpose of the cleanup

defer is used to execute the code after the function exection completed

if there are multiple defere statement in a funnction call then it follow lifo 

*/