package main

import "fmt"
import "testing"

   
func Benchmark_Day(t *testing.B) {
    fmt.Println("Manager Start Work")
    b := 0
    r := 8
	for numWorkers := 1; numWorkers <= r; numWorkers *= 2 {
		b++
		t.Run(fmt.Sprintf("Builder#%d Group", b), func(t *testing.B) {
			
		Day(r)

		})
		s := Day(r)
		fmt.Println("Builder work " , s)	
	}
	
	
}
