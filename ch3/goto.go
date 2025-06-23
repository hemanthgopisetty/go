package main

import (
	"crypto/rand"
	"fmt"
)

func gotoS() {
	a:=rand.Intn(10)
	for a<100{
		if a%5 ==0 {
			goto done
		}
		a = a*2 +1 
	}
	done :
	  fmt.Println("Do Complicated Stuff")
}

/*

Dijkstra Said Goto is very harmful

goto Traditionally can jump nearly to anywhere

In go goto statement execution jumps to it

Go put some restriction over it

	a := 20
	goto skip
	b := 30
skip:
	c := 30
	if c > a{
		goto inner
	}
	inner :
		fmt.Print("Confusing")


don't use goto staement to get more confused
again
use goto for greater purpose

like for that we need a boleean 
state as a flag , we should
maintain 

at this point we can use go to




*/
