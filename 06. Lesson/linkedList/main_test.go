package main

import (
"testing"
"fmt"
)

func TestList(t *testing.T) {
fmt.Println("Testing ListNode, with 4")
pos:= 1

node3 := ListNode {
value: 0, next: nil,}

tail := ListNode {
value: -4, next: &node3,}

node3 = ListNode {
value: 0, next: &tail,}

node2 := ListNode {
value: 2, next: &node3,}

if node2.value == 0 {}

head := ListNode {
value: 3, next: &node2,}

fmt.Println(hasCycle(&head, pos))
fmt.Println()
}

func TestList2(t *testing.T) {
fmt.Println("Testing ListNode, with 2")


head := ListNode {
value: 0, next: nil,}

pos:= 0

tail := ListNode {
value: 2, next: &head,}

head = ListNode {
value: 1, next: &tail,}

fmt.Println(hasCycle(&head, pos))
fmt.Println()
}

func TestList3(t *testing.T) {
fmt.Println("Testing ListNode, with 1")

pos:= -1

head := ListNode {
value: 1, next: nil,}

fmt.Println(hasCycle(&head, pos))
fmt.Println()
}

func TestList4(t *testing.T) {
fmt.Println("Testing ListNode, nil adress in tail")
pos:= 1

node3 := ListNode {
value: 0, next: nil,}

tail := ListNode {
value: -4, next: nil,}

node3 = ListNode {
value: 0, next: &tail,}

node2 := ListNode {
value: 2, next: &node3,}

if node2.value == 0 {}

head := ListNode {
value: 3, next: &node2,}

fmt.Println(hasCycle(&head, pos))
fmt.Println()
}
