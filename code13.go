package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 10000 {
			break
		}
		fmt.Println(x)
		x++
	}

	//for i:=1;i<=10000; i++{
	//	fmt.Println(i)
	//}
}
