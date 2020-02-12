package main

import (
	"fmt"
)


func main(){
	str := []string{"a","a","b","b","d","b","c","c"}
}

func RemoveDuplicates(str []string)[]string{
	for i := 0;i<len(str);i++{
		if str[i] == str[i+1]{
			copy(str[i:],str[i+1:])
			str = str[:len(str)-1]
			i--
		}
	}
	return str
}


