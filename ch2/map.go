package main



func practicemap() {
	
}

/*

built in data type for situtations where want to assosicate the value to another

var declaration 
var nilMap map[string]int;

zero value for map is nil
nil map has length of zero 

read nil map returns zero value for the map values type
attempting to write a nil map causes a panic

: = 
declaration via 
:= map[string]int

	var mp map[string]int;
	mp2:= map[string]int{};
	fmt.Println(len(mp),len(mp2));


		mp2:= map[string]int{
		"raju" : 5,
		"gaju" : 6,
	};

	for empty map literal we can read and write
	but for var declaration reading gives 0 and writing causes 
	panic error

	we can also create map using make


	maps are like slices , they grow ,they are not comparable
	there is some restriction in map key type

	use slice for continous data
	map for key value

	in go map is hash map
	go includes a hash implementation as part of runtime


	read map value like this 

	mp3["raju"]
	mp3["raju"]=6

	map returns value zero if that keys is not in the map

	Comma idiom
	var mp map[string]int;
	x,y := mp["3"]
	fmt.Println(x,y)
    first get value of the key
	second gets the boolean wether the key is in map or not ?

	you can delete the key and value using delete the func
	
    delete(mp,"2")''

	empty map using 
	clear function 
	clear(mpa)

	comparing maps using built in functions 

	maps.EqualFunc()



Using maps as sets

set data type ensures unique ness among elements 
we can use map as set itself
intSet:=map[int]bool{}

we can also do the same thing with strucuts
by the way structs take zero byte where as boolean takes one bytes

 */
