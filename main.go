package main

import (
	"fmt"
	"github.com/sgrumley/DataStructures/binaryTree"
	"github.com/sgrumley/DataStructures/disjointSet"
	"github.com/sgrumley/DataStructures/linkedList"
	"github.com/sgrumley/DataStructures/queue"
	"github.com/sgrumley/DataStructures/set"
	"github.com/sgrumley/DataStructures/stack"
)

func main() {
	fmt.Println("TODO:")
	fmt.Println("Skip lists, Hashtables, Treap, Fibonacci Heap, Huffman Tree, Splay Tree, Tries(compressed, x-fast), Suffix Tree, multi dimensional Search tree")
	fmt.Println("DONE:")
	fmt.Println("Disjoints Sets, Linked List, Set")

	fmt.Println("**************************** SET ****************************")
	nameSet := set.New()
	nameSet.Add("sam")
	nameSet.Add("bob")
	nameSet.List()
	fmt.Println(nameSet.Has("sam"))
	fmt.Println(nameSet.Has("jake"))

	fmt.Println("**************************** STACK ****************************")
	nameStack := stack.New()
	nameStack.Push("sam")
	nameStack.Push("bob")
	fmt.Println(nameStack.Top())
	name1, _ := nameStack.Pop()
	fmt.Println(name1)
	fmt.Println(nameStack.IsEmpty())
	name2, _ := nameStack.Pop()
	fmt.Println(name2)
	fmt.Println(nameStack.IsEmpty())
	_, operationComplete := nameStack.Pop()
	fmt.Println(operationComplete)

	fmt.Println("**************************** QUEUE ****************************")
	nameQueue := queue.New()
	nameQueue.Push("sam")
	nameQueue.Push("bob")
	fmt.Println(nameQueue.Top())
	name1q, _ := nameQueue.Pop()
	fmt.Println(name1q)
	fmt.Println(nameQueue.IsEmpty())
	name2q, _ := nameQueue.Pop()
	fmt.Println(name2q)
	fmt.Println(nameQueue.IsEmpty())
	_, operationCompleteq := nameQueue.Pop()
	fmt.Println(operationCompleteq)

	fmt.Println("**************************** DISJOINT SET ****************************")
	d := disjointSet.New()
	d.Init(10)
	d.PrintSet()
	d.UnionSet(3, 7)
	d.UnionSet(7, 9)
	fmt.Println(d.Find(3))
	d.PrintSet()

	fmt.Println("**************************** LINKED LIST ****************************")
	fmt.Println("Work in progress")
	linkedList := linkedList.New()
	linkedList.Insert(1)
	linkedList.Insert(4)
	linkedList.Insert(7)
	linkedList.Insert(10)
	linkedList.Insert(21)
	linkedList.Print()
	linkedList.Delete(10)
	linkedList.Print()

	fmt.Println("**************************** SKIP LIST ****************************")
	fmt.Println("Work in progress")

	fmt.Println("**************************** BINARY TREE ****************************")
	bTree := binaryTree.New()
	bTree.Add(87)
	bTree.Add(6)
	bTree.Add(234)
	bTree.Add(35)
	bTree.Add(136)
	bTree.Add(46)
	bTree.List()

	fmt.Println("**************************** BINARY TREE SLICE ****************************")
	bTreeS := binaryTree.NewBTSlice(15)
	bTreeS.AddRoot("A")
	bTreeS.AddLeft("B", 0)
	bTreeS.AddRight("C", 0)
	bTreeS.AddLeft("D", 1)
	bTreeS.AddLeft("E", 2)
	bTreeS.AddLeft("F", 3)
	bTreeS.AddLeft("G", 5)
	bTreeS.List()

}
