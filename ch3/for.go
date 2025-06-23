package main

import "fmt"

func forLoop() {
	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i)
	// }
	// i := 0
	// for i < 100 {
	// 	fmt.Println(i)
	// 	i *= 3
	// }

	// even := []int{1, 2, 3, 4, 5, 6}
	// names := map[string]int{}
	// names["Raju"]++
	// names["Jaya"]++
	// for i, v := range even {
	// 	fmt.Println(i, v)
	// }
	// for i, v := range names {
	// 	fmt.Println(i, v)
	// }

	samples := []string{"hello", "apple_n!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r)
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
	
}

/*

there are four ways of writing for loop

1.c style
  for i:=0;i<=5;i++{
  	fmt.Println(i)
  }
  no need of parantheses
  around it;s parts like if and else

  above for loop has three
  parts
   1.intiliazation(sets one or more variables) before loop begins
  	use := to intialize
  	not var
  	we can shadow the existing variables in this

   2.Comparision must be an expression
  	checked immediately after intialization
  	checked before the iteration of loop
  	if cond is true then loop iterates
   3.loop variable math (increment,decrement,division)

  we can omit any of the part
  make sure not to infinite the loop

2.condition only
	When we leave off both the intialization
	and incrementation
	just only condition
	like
    i := 0
	for i < 100 {
		fmt.Println(i)
		i *= 3
	}

3.infinite
	omit the condtion and intialization
	incrementation

	then this type of loop is called for loop

4.for range
	THis loop is to iterate over the
	built in types of go
	and compound types of go

	even := []int{1, 2, 3, 4, 5, 6}
	names := map[string]int{}
	names["Raju"]++
	names["Jaya"]++
	for i, v := range even {
		fmt.Println(i, v)
	}
	for i, v := range names {
		fmt.Println(i,v)
	}

	for range copy the values dude
	not a reference


To get out of
	loop -> break
	iteration -> continue

Labeling the For Loop Statement

	samples := []string{"hello", "apple_n!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r)
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}

	in some cases this loop also useful



Choose Right For Statement
We can use any for statement
mostly range and completed for loop i used
and with combined break and continue




In go There is do statement


*/
