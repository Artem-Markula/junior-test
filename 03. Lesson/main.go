package main
 
import (
    "time"
    "math/rand"	
)
func main() {}
 	
func Day(Count int)(W int) {
    W = 0		
    c := make(chan int) 
    for i := 0; i < Count; i++ {
        go Worker(i, c)
	
    }
   
    W = <-c	
    return W
}
 
func Worker(id int, c chan int) { 
    rand.Seed(time.Now().UnixNano())
    work := time.Duration(rand.Intn(3))
		
    time.Sleep(work * time.Second)
    c <- id 
}	
    
