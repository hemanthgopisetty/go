package main

import "fmt"

func switchS() {
	words := []string{
		"a",
		"cow",
		"smile",
		"gopher",
	}
	for _, word := range words {
		switch size:=len(word);size{
		case 1,2,3,4:
			fmt.Println(word,"is a short word")
		case 5 :
			fmt.Println(word,"is excatly the right lenght")
		case 6,7,8,9:
		default:
			fmt.Println(word,"The word is too long")
		}
		
	}
}

/*
like many c-derived languages
go has switch statement

but go's switch is different from
other langauges , and it makes the switch more useful
that other

as in if we can declare a varaible
in switch
for switch scope




*/
