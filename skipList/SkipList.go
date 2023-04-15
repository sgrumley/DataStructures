// TODO intergrate search to be used by delete and insert
package skipList

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_LEVEL = 10

type Node struct {
	key     int
	forward []*Node
}

func nodeBuilder(lvl int, key int) *Node {
	return &Node{
		key:     key,
		forward: make([]*Node, lvl, lvl),
	}
}

type SkipList struct {
	lvl  int
	head *Node
}

func SkipListBuilder() *SkipList {
	s := new(SkipList)
	s.lvl = 0
	s.head = nodeBuilder(MAX_LEVEL, 0)

	return s
}

func randomLevel() int {
	lvl := 1

	for {
		rand.Seed(time.Now().UnixNano())
		if rand.Intn(2) == 0 || lvl > MAX_LEVEL {
			break
		}
		lvl++
	}
	fmt.Println("The randomm level lvl is : ", lvl)
	return lvl
}

func (sl *SkipList) insert(key int) bool {
	var temp, new *Node
	update := make([]*Node, MAX_LEVEL, MAX_LEVEL)
	temp = sl.head
	k := sl.lvl

	// Search
	for i := k; i >= 0; i-- {
		for temp != nil && temp.forward[i] != nil && temp.forward[i].key < key {
			temp = temp.forward[i]
		}
		update[i] = temp
	}

	if temp != nil && temp.key == key {
		return false
	}

	k = randomLevel()
	fmt.Println(key, ":", k)

	//update head  pointer
	if k > sl.lvl {
		for i := sl.lvl; i < k; i++ {
			update[i] = sl.head
		}
		sl.lvl = k
	}

	new = nodeBuilder(k, key)

	for i := 1; i < k; i++ {
		if update[i] != nil {
			new.forward[i] = update[i].forward[i]
			update[i].forward[i] = new
		}
	}

	return true
}

func (sl *SkipList) delete(key int) bool {
	update := make([]*Node, MAX_LEVEL, MAX_LEVEL)
	temp := sl.head

	for i := sl.lvl - 1; i >= 0; i-- {
		for temp != nil && temp.forward[i] != nil && temp.forward[i].key < key {
			temp = temp.forward[i]
		}
		update[i] = temp
	}

	if temp.forward != nil && temp.forward[0].key != key {

		return false
	}

	deleteN := temp.forward[0]
	for i := sl.lvl - 1; i >= 0; i-- {
		if update[i] != nil && update[i].forward[i] == deleteN {
			update[i].forward[i] = deleteN.forward[i]
			if sl.head.forward[i] == nil {
				sl.lvl--
			}
		}
	}
	return true
}

func (sl *SkipList) search(key int) int {
	var temp *Node
	temp = sl.head
	for i := sl.lvl - 1; i >= 0; i-- {
		for temp != nil && temp.forward[i] != nil && temp.forward[i].key <= key {
			if temp.forward[i].key == key {
				return temp.forward[i].key
			}
			temp = temp.forward[i]
		}
	}
	return -1
}

func (sl *SkipList) print() {
	var temp *Node
	temp = sl.head
	fmt.Println("Max levels: ", MAX_LEVEL)
	fmt.Println("Highest Level: ", sl.lvl)
	fmt.Println("Key:Next")

	for i := sl.lvl; i >= 0; i-- {
		fmt.Println("\nLevel:", i)
		if temp != nil {
			//fmt.Print(temp.key, " -> ")
		}
		for temp != nil { //&& temp.forward[i] != nil {
			if temp.forward[i] != nil {
				fmt.Print(temp.forward[i].key, " -> ")
			} else {
				//fmt.Print(temp.key, " -> ")
				break
			}

			temp = temp.forward[i]
		}
	}
	fmt.Println(" ")
}

func main() {

	sl := SkipListBuilder()

	// for i := 1; i < 10; i++ {
	// 	sl.insert(i)
	// }
	sl.insert(1)

	sl.insert(2)

	sl.insert(3)
	sl.print()

	// fmt.Println(sl.search(3))
	// fmt.Println(sl.delete(3))
	// fmt.Println(sl.search(3))
}
