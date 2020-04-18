package main

import "fmt"
import "testing"

func TestIsomorphic(t *testing.T) {
	fmt.Println("Test: Isomorphic 3-Words")

	c := "foo"	

	s := "bar"

	fmt.Println(isIsomorphic(c, s))

	c = "egg"	

	s = "add"

	fmt.Println(isIsomorphic(c, s))
	fmt.Println()
}

func TestIsomorphic2(t *testing.T) {
	fmt.Println("Test: 5-Words")

	c := "paper"	

	s := "title"

	fmt.Println(isIsomorphic(c, s))
	
	c = "happy"	

	s = "sadly"

	fmt.Println(isIsomorphic(c, s))

	fmt.Println()
}

func TestIsomorphic3(t *testing.T) {
	fmt.Println("Test: Different length")

	c := "sdfd"	

	s := "a"

	fmt.Println(isIsomorphic(c, s))
	fmt.Println()
}

func TestIsomorphic4(t *testing.T) {
	fmt.Println("Test: Simple Isomorphic")

	c := "bb"	

	s := "aa"

	fmt.Println(isIsomorphic(c, s))
	fmt.Println()
}

func TestIsomorphic5(t *testing.T) {
	fmt.Println("Test: Midle Isomorphic")

	c := "abbbtttcc"	

	s := "abbbaaacc"

	fmt.Println(isIsomorphic(c, s))
	fmt.Println()
	
}

func TestIsomorphic6(t *testing.T) {
	fmt.Println("Test: Null String")
 
	c := " "	

	s := " "

	fmt.Println(isIsomorphic(c, s))
	fmt.Println()

}
