package day18

import (
	_ "embed"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

type Node struct {
	left interface{}
	right interface{}
	up *Node
}

func parseNodes(s string) *Node {
	res := &Node{}
	path := make([]*Node, 0)
	isLeft := true

	path = append(path, res)

	for _, c := range strings.Split(s[1:], "") {
		switch c {
		case "[":
			newNode := &Node{ up: path[len(path)-1] }
			if isLeft {
				path[len(path)-1].left = newNode
			} else {
				path[len(path)-1].right = newNode
			}

			path = append(path, newNode)
			isLeft = true
		case "]":
			path = path[:len(path)-1]
		case ",":
			isLeft = false
		default:
			val, _ := strconv.Atoi(c)
			if isLeft {
				path[len(path)-1].left = val
			} else {
				path[len(path)-1].right = val
			}
		}
	}
	return res
}

func sumTrees(tree1, tree2 *Node) *Node {
	return &Node{
		left: tree1,
		right: tree2,
		up: nil,
	}
}

func isChild(node *Node) bool {
	return node.up != nil
}

func isLeftChild(node *Node) bool {
	switch node.up.left.(type) {
	case *Node:
		return node.up.left.(*Node) == node
	default:
		return false
	}
}

func isRightChild(node *Node) bool {
	switch node.up.right.(type) {
	case *Node:
		return node.up.right.(*Node) == node
	default:
		return false
	}
}

func reduceTree(tree *Node) {

	for canContinue := true; canContinue; {

		canContinue = false

		// Find a node of two ints, with depths >= 4
		nodeToExplode := findNodeToExplode(tree, 0)
		if nodeToExplode != nil {

			fmt.Println("----- Exploding node: -----")
			printTree(nodeToExplode, "EE ")

			incrementUpperLeftInt(nodeToExplode, nodeToExplode.left.(int), false)
			incrementUpperRightInt(nodeToExplode, nodeToExplode.right.(int), false)

			// Remove exploded node
			if isLeftChild(nodeToExplode) {
				nodeToExplode.up.left = 0
			} else {
				nodeToExplode.up.right = 0
			}

			fmt.Println("----- After explode -----")
			printTree(tree, "")

			canContinue = true
			continue
		}

		canContinue = splitNodes(tree)
		if canContinue {
			fmt.Println("----- After split -----")
			printTree(tree, "")
		}
	}

}

func incrementUpperLeftInt(node *Node, value int, mustGoDown bool) bool {
	if ! mustGoDown {
		if ! isChild(node) {
			// Nothing above - failure!
			return false
		}
		if isRightChild(node) {
			// Is right child: go down ASAP on the left side, will always succeed
			switch node.up.left.(type) {
			case int:
				node.up.left = node.up.left.(int) + value
				return true

			default:
				return incrementUpperLeftInt(node.up.left.(*Node), value, true)
			}

		} else {
			// Is left child: must continue upwards
			return incrementUpperLeftInt(node.up, value, false)
		}
	}

	// Must go down, always on the right!
	switch node.right.(type) {
	case int:
		node.right = node.right.(int) + value
		return true
	default:
		return incrementUpperLeftInt(node.right.(*Node), value, true)
	}

	// Unreachable?
	return false
}

func incrementUpperRightInt(node *Node, value int, mustGoDown bool) bool {
	if ! mustGoDown {
		if ! isChild(node) {
			// Nothing above - failure!
			return false
		}
		if isLeftChild(node) {
			// Is left child: go down ASAP on the right side, will always succeed
			switch node.up.right.(type) {
			case int:
				node.up.right = node.up.right.(int) + value
				return true

			default:
				return incrementUpperRightInt(node.up.right.(*Node), value, true)
			}

		} else {
			// Is right child: must continue upwards
			return incrementUpperRightInt(node.up, value, false)
		}
	}

	// Must go down, always on the left!
	switch node.left.(type) {
	case int:
		node.left = node.left.(int) + value
		return true
	default:
		return incrementUpperRightInt(node.left.(*Node), value, true)
	}

	// Unreachable?
	return false
}

func findNodeToExplode(tree *Node, depth int) *Node {
	if reflect.TypeOf(tree.left).Kind() == reflect.Int && reflect.TypeOf(tree.right).Kind() == reflect.Int && depth > 4 {
		return tree
	}

	switch tree.left.(type) {
	case *Node:
		res := findNodeToExplode(tree.left.(*Node), depth + 1)
		if res != nil {
			return res
		}
	}

	switch tree.right.(type) {
	case *Node:
		res := findNodeToExplode(tree.right.(*Node), depth + 1)
		if res != nil {
			return res
		}
	}

	return nil
}

func splitNodes(node *Node) bool {
	switch node.left.(type) {
	case int:
		value := node.left.(int)
		if value >= 10 {
			node.left = &Node{
				left:  int(math.Floor(float64(value) / 2.0)),
				right: int(math.Ceil(float64(value) / 2.0)),
				up:    node,
			}
			return true
		}
	default:
		if splitNodes(node.left.(*Node)) {
			return true
		}
	}

	switch node.right.(type) {
	case int:
		value := node.right.(int)
		if value >= 10 {
			node.right = &Node{
				left:  int(math.Floor(float64(value) / 2.0)),
				right: int(math.Ceil(float64(value) / 2.0)),
				up:    node,
			}
			return true
		}
	default:
		if splitNodes(node.right.(*Node)) {
			return true
		}
	}

	// Unreachable?
	return false
}

func printTree(node *Node, indent string) {
	switch node.left.(type) {
	case *Node:
		fmt.Printf("%sLeft:\n", indent)
		printTree(node.left.(*Node), indent + ".")
	case nil:
		fmt.Printf("%sLeft: NIL\n", indent)
	default:
		fmt.Printf("%sLeft: %d\n", indent, node.left.(int))
	}

	switch node.right.(type) {
	case *Node:
		fmt.Printf("%sRight:\n", indent)
		printTree(node.right.(*Node), indent + ".")
	case nil:
		fmt.Printf("%sRight: NIL\n", indent)
	default:
		fmt.Printf("%sRight: %d\n", indent, node.right.(int))
	}

}

func Solve() {

	tree := parseNodes("[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]")

	for _, s := range(strings.Split(data, "\n")) {
		tree = sumTrees(tree, parseNodes(s))
		reduceTree(tree)
	}

}