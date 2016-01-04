package main

import "fmt"

func add(x int, y int) int{
	return x+y
}

func div(x int, y int) int{
	return x/y
}

func main(){
	fmt.Println(add(3,2))
	fmt.Println(div(1,0))
}


