package main

import "fmt"
import "testing"

func TestCheck(t *testing.T) {
	fmt.Println("Test: Parenthesis")

	f := "()"	

	fmt.Println(check(f))
	fmt.Println()
}


func TestCheck2(t *testing.T) {
	fmt.Println("Test: All brackets")

	f := "()[]{}"	

	fmt.Println(check(f))
	fmt.Println()
}

func TestCheck3(t *testing.T) {
	fmt.Println("Test: 2 Brackets")

	f := "{[]}"	

	fmt.Println(check(f))
	fmt.Println()
}

func TestCheck4(t *testing.T) {
	fmt.Println("Fail brackets")

	f := "(]"	

	fmt.Println(check(f))
	fmt.Println()

	f = "([)]"	

	fmt.Println(check(f))
	fmt.Println()
}

func TestCheck5(t *testing.T) {
	fmt.Println("bracket with words")

	f := "r4*{3d-[f2%]s}1f"	

	fmt.Println(check(f))
	fmt.Println()
}

func TestCheck6(t *testing.T) {
	fmt.Println("Null string")

	f := " "	

	fmt.Println(check(f))
	fmt.Println()
}
