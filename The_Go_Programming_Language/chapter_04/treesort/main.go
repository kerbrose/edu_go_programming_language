/*
TODO: write a stack to demonstrates non-recursive of of tree-traversing (DFS)
*/

package main

import "fmt"

type tree struct {
	value       int
	left, right *tree
}

type stack []*tree

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t // Equivalent to return &tree{value: value}
	}

	if value > t.value {
		t.right = add(t.right, value)
	} else {
		t.left = add(t.left, value)
	}

	return t
}

func inOrder_DFS_Recursive(t *tree) {
	if t == nil {
		return
	}

	inOrder_DFS_Recursive(t.left)
	fmt.Print(t.value, " ")
	inOrder_DFS_Recursive(t.right)
}

func push(s stack, t *tree) {
	s = append(s, t)
}

func pop(s stack) {

}

func inOrder_DFS_Stack(t *tree) {

	var s stack

	if t == nil {
		return
	}

	// push
	push(s, t)

}

/*

fmt.Println(add(&root, -1))
fmt.Println(add(&root, 2))
fmt.Println(add(&root, 5))
fmt.Println(add(&root, 2))
fmt.Println(add(&root, 4))

  	 0
	-1   2
			2	 5
				4

*/
func main() {
	var root tree

	inOrder_DFS_Recursive(&root)
	//fmt.Println()
	//fmt.Println(add(&root, 2))
}
