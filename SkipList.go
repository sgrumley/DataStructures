package main

type SNode struct {
	before *Node
	after *Node
	above *Node
	below *Node
	value interface{}
}

type SkipList struct {
	head *Node
	tail *Node
	size int
}

func (s *SkipList) Insert(value interface{}){

}

func (s *SkipList) Delete(value interface{}){

}

func (s *SkipList) Find(value interface{}){

}

func (s *SkipList) Print(value interface{}){

}

func main() {

}
