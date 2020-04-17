package main

import "fmt"
import "testing"

func TestNumbers(t *testing.T) {
	fmt.Println("Test: Regular Numbers")

	target := 9	

	nums := []int{2,2,4,5}

	fmt.Println(twoSum(nums, target))
}

func TestNumbersNone(t *testing.T) {
	fmt.Println("Test: Array, Don't have Numbers")

	target := 9	

	nums := []int{1,5,1,2}

	fmt.Println(twoSum(nums, target))
}

func TestNull(t *testing.T) {
	fmt.Println("Test: Null numbers")

	target := 9	

	nums := []int{}
	

	fmt.Println(twoSum(nums, target))
}
