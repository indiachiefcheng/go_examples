package main

import (
	"fmt"
)

func main(){
	var x,y []int
	for i := 0;i<10;i++{
		y = append(x,i)
		fmt.Printf("%d cap=%d\t\n",i,cap(y),y)
		x = y
	}
}
