package binaryTree

import "fmt"

type BinaryTreeSlice []string

func NewBTSlice(size int) BinaryTreeSlice {
	return make(BinaryTreeSlice, size)
}

func (bts *BinaryTreeSlice) AddRoot(key string) bool {
	if (*bts)[0] != "" {
		return false
	}

	(*bts)[0] = key
	return true
}

func (bts *BinaryTreeSlice) AddLeft(key string, parent int) bool {
	if (*bts)[parent] == "" {
		return false
	}
	fmt.Println(key)
	(*bts)[(parent*2)+1] = key

	return true
}

func (bts *BinaryTreeSlice) AddRight(key string, parent int) bool {
	if (*bts)[parent] == "" {
		return false
	}
	fmt.Println(key)
	(*bts)[(parent*2)+2] = key

	return true
}

func (bts *BinaryTreeSlice) List() {
	for _, val := range *bts {
		if val == "" {
			fmt.Printf("-")
		} else {
			fmt.Printf("%s", val)
		}
	}
}
