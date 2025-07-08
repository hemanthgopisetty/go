package main


type Person struct{
	Firstname string
	Lastname string
	Age int 
}
func makeMult(base int) func(int) int {
 return func(factor int) int {
 return base * factor
 }
}
func closures(){
	// a:=20
	// f:=func(){
	// 	a=30
	// 	a:=20
	// 	fmt.Println(a)
	// }
	// f()
	// fmt.Println(a)
	// people:=[]Person{
	// 	{"Pat", "Patterson", 37},
 	// 	{"Tracy", "Bobdaughter", 23},
 	// 	{"Fred", "Fredson", 18},
	// }
	// fmt.Println(people)
	// sort.Slice(people,func(i,j int)bool{
	// 	return people[i].Lastname < people[j].Lastname;
	// });

}

/*

functions declared inside functions are special

they are CLOSURES

functions declared inside functions are able to access
the variables and modify the variables in 
outer function

Example
	a:=20
	f:=func(){
		a=30
	}
	f()
	fmt.Println(a)

we can also shadow the variable too 

	a:=20
	f:=func(){
		a=30
		a:=20
		fmt.Println(a)
	}
	f()
	fmt.Println(a)

	closures allow you to do a limit the function's scope

	closures are really intresting when they are passed to other functions 
	or returned from other functions 
	they allow you to take the variable within your function 
	and use those values outside the function of your function 

	Passing function as Parameters
	we can pass function as paramters since functions are values

	closure a very useful pattern used in many standar libraries
	people:=[]Person{
		{"Pat", "Patterson", 37},
 		{"Tracy", "Bobdaughter", 23},
 		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	sort.Slice(people,func(i,j int)bool{
		return people[i].Lastname < people[j].Lastname;
	});
	if we see people is accesible in the passed fucntion becuase 
	of clousre
	the known a person is captured by function 


	Returning function from another function

	in addition to passing we can also return the function from an function
	func makeMult(base int) func(int) int {
 return func(factor int) int {
 return base * factor
 }
	twoBase := makeMult(2)
 threeBase := makeMult(3)
 for i := 0; i < 3; i++ {
 fmt.Println(twoBase(i), threeBase(i))
 }


*/