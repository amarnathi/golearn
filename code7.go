package main

import "fmt"

func main() {

	x7:=19
	y7:=21

	eq:=x7==y7
	lteq:=x7<=y7
	gteq:=x7>=y7
	nteq:=x7!=y7
	lt:=x7<y7
	gt:=x7>y7

	fmt.Printf("%T\t%v\n",eq,eq)
	fmt.Printf("%T\t%v\n",lteq,lteq)
	fmt.Printf("%T\t%v\n",gteq,gteq)
	fmt.Printf("%T\t%v\n",nteq,nteq)
	fmt.Printf("%T\t%v\n",lt,lt)
	fmt.Printf("%T\t%v\n",gt,gt)
}
