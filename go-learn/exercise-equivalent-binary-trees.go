package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {

	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		x := <- ch1
		y := <- ch2
		fmt.Println(x, y)
		if x != y {
			return false
		}
	}

	return true
}

func main() {
	flag1 := Same(tree.New(1), tree.New(1))
	flag2 := Same(tree.New(1), tree.New(4))

	fmt.Println(flag1, flag2)
}

