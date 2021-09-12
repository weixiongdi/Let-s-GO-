//package Let_s_GO_


package main

import (
	"fmt"
)

func chvalue( a int) int {
	a=a+1
	return a
}

func chpointer (a *int){
	*a =*a+1
	return
}
func main(){
	a:=10
	b:=chvalue(a)
	fmt.Println(b)
	fmt.Println(a)

	chpointer(&a)
	fmt.Println(a)
}