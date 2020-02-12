package main

import (
	"fmt"
)

func main(){
	slice := []int{1,2,3,4,5}
	fmt.Println("original slice is:",slice)
	fmt.Println("the rotated slice is:",rotate(slice))
	
}

func rotate(slice []int) []int{
	for i,j := 0,len(slice)-1;i<j;i,j = i+1,j-1{
		slice[i],slice[j] = slice[j],slice[i]
	}
	return slice
}
