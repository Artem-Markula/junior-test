package main

import s "strings"
import "fmt"

func isIsomorphic(c string, t string) (answer bool) {
	join := 0
	joined := 0
	var num int  = 0 
	map2:= map[string] int{
	"q":1, "w":1, "e":1, "r":1, "t":1, "y":1, "u":1, "i":1, "o":1, "p":1,
	"a":1, "s":1, "d":1, "f":1, "g":1, "h":1, "j":1, "k":1, "l":1, "z":1,
	"x":1, "c":1, "v":1, "b":1, "n":1, "m":1,
	}
	
	arr:= []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
	"a", "s", "d", "f", "g", "h", "j", "k", "l", "z",
	"x", "c", "v", "b", "n", "m"}
	
	el1 := len(c)
	el2 := len(t)
	
	for z:= 0; z < len(arr); z++{
	join = join + s.Count(c, arr[z])
	joined = joined + s.Count(t, arr[z])	
	}
	
	if join == 0{
	fmt.Println("String c dont have words")
	answer = false
	}
	
	if joined == 0{
	fmt.Println("String s dont have words")
	answer = false
	}
	

	if el1 == el2 {
	
	for i:= 0; i<len(arr); i++{
	
	number := s.Count(t, arr[i])
	
	if number >= 2{
	num = map2[arr[i]]
	}
	
	}
	for i:= 0; i<len(arr); i++{
	
	number := s.Count(c, arr[i])
	
	if number >= 2{
	num = num + map2[arr[i]]
	}
	
	}
	}else{
	answer = false
	}
	
	if num >= 2{
	answer = true
	}

	return answer

}


func main() {
	
	
}
