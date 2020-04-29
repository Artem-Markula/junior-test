package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func hasCycle(head *ListNode, pos int) bool {

  if head.next == nil{
  fmt.Println("There is no cycle in the linked list")	
  return false	
}
  	

  node1 := head.next
  node2:= node1.next
  node3:= node2.next
  node4:= node3.next
  
  value1:= head.value
  value2:= node1.value
  value3:= node2.value
  value4:= node3.value

  data := make([]int, 4)
  data[0] = value1
  data[1] = value2
  data[2] = value3
  data[3] = value4

  if node4 != nil{
  fmt.Println("There is a cycle in the linked list, where tail connects to the" , data[pos], "node")
  return true
  }else{
  fmt.Println("There is no cycle in the linked list")
  return false
  }

  if pos < 0 {
  fmt.Println("There is no cycle in the linked list")
  
  return false
  	
  }

  if pos >= len(data) {
  fmt.Println("Position out of range in linked list")
  return false
  	
  }
			
  return true
}


func main(){

}
