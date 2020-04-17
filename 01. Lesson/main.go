package main

import "fmt"


func twoSum(nums []int, target int)(f string){
	a1 := 0
	a2 := 0

	sum:= [50]int{}
	for j:= 1; j < len(nums); j++ {
	for i:=0; i<len(nums); i++{
	sum[i] = nums[i] + nums[j]
	
	

	if sum[i] == target {
	a1 = i
	a2 = j
	fmt.Println("Sum =", target, "with two elements in array, with index",a1, a2)
	return ""
		
}		
	
}

}	
	if len(nums) == 0 {
	fmt.Println("Array dont have numbers")
	return ""
        }

	if a1 + a2 != target {
	fmt.Println("Array dont have sum two elements =", target)
	return ""
        }
		

return "Error"
}

func main() {
	
}
