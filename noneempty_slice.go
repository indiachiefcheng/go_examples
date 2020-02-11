package main

import "fmt"

func noneempty(strings []string)[]string{
	i := 0
	for _,s := range strings{
		if s != ""{
			strings[i] =s
			i++
		}
	}
	return strings[:i]
}

func noneempty2(strings []string)[]string{
	out := strings[:0] //zero-length slice of original
	for _,s := range strings{
		if s != ""{
			out = append(out,s)
		}
	}
	return out
}
