package day12

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed data.txt
var data string

type Node struct {
	name string
	isLarge bool
	neighbors []*Node
}

func canGoThere(rootPath []string, node *Node) bool {
	if node.isLarge { return true }
	if node.name == "start" { return false }

	for _, previousNode := range rootPath {
		if previousNode == node.name {
			return false
		}
	}
	return true
}

func step1(start *Node, end *Node) {
	rootPath := make([]string, 1)
	rootPath[0] = start.name

	allPaths := iterateOnPaths(start, rootPath, end)
	//for _, path := range allPaths {
	//	fmt.Printf("%v\n", path)
	//}
	fmt.Println(len(allPaths))
}

func iterateOnPaths(start *Node, rootPath []string, end *Node) [][]string {
	paths := make([][]string, 0)
	//fmt.Printf("Iterating from %v\n", rootPath)

	for _, nextNode := range start.neighbors {

		if ! canGoThere(rootPath, nextNode) {
			//fmt.Printf("FORBIDDEN: %v -> %s\n", rootPath, nextNode.name)
			continue
		}

		nextRootPath := make([]string, len(rootPath) + 1)
		copy(nextRootPath, rootPath)
		nextRootPath[len(nextRootPath) - 1] = nextNode.name

		if nextNode == end {
			//fmt.Printf("ENDING ON: %v -> %s\n", rootPath, nextNode.name)
			paths = append(paths, nextRootPath)
			continue
		} else {
			paths = append(paths, iterateOnPaths(nextNode, nextRootPath, end)...)
		}
	}

	return paths
}

func step2(start *Node, end *Node) {
	rootPath := make([]string, 1)
	rootPath[0] = start.name

	allPaths := iterateOnPathsWithJoker(start, rootPath, end, true)
	fmt.Println(len(allPaths))
}

func iterateOnPathsWithJoker(start *Node, rootPath []string, end *Node, joker bool) [][]string {

	paths := make([][]string, 0)
	//fmt.Printf("========= FROM %v j=%v\n", rootPath, joker)

	for _, nextNode := range start.neighbors {
		nextJoker := joker

		if ! canGoThere(rootPath, nextNode) {
			if nextJoker && nextNode.name != "start" {
				//fmt.Printf("JOKER: %v -> %s\n", rootPath, nextNode.name)
				nextJoker = false
			} else {
				//fmt.Printf("FORBIDDEN: %v -> %s, j=%v\n", rootPath, nextNode.name, nextJoker)
				continue
			}
		}

		nextRootPath := make([]string, len(rootPath) + 1)
		copy(nextRootPath, rootPath)
		nextRootPath[len(nextRootPath) - 1] = nextNode.name

		if nextNode == end {
			//fmt.Printf("ENDING ON: %v -> %s\n", rootPath, nextNode.name)
			paths = append(paths, nextRootPath)
			continue
		} else {
			paths = append(paths, iterateOnPathsWithJoker(nextNode, nextRootPath, end, nextJoker)...)
		}

	}
	//fmt.Println("^^^^^^^^^")

	return paths
}

func isLargeCave(s string) bool {
	return strings.ToUpper(s) == s
}

func Solve() {
	startNode := Node{"start", isLargeCave("start"), make([]*Node, 0)}

	allNodes := make(map[string]*Node)
	allNodes["start"] = &startNode

	for _, line := range strings.Split(data, "\n") {
		endpoints := strings.Split(line, "-")

		ep0, ep0Exists := allNodes[endpoints[0]]
		if ! ep0Exists {
			ep0 = &Node{endpoints[0], isLargeCave(endpoints[0]), make([]*Node, 0)}
			allNodes[endpoints[0]] = ep0
		}

		ep1, ep1Exists := allNodes[endpoints[1]]
		if ! ep1Exists {
			ep1 = &Node{endpoints[1], isLargeCave(endpoints[1]),make([]*Node, 0)}
			allNodes[endpoints[1]] = ep1
		}

		ep0.neighbors = append(ep0.neighbors, ep1)
		ep1.neighbors = append(ep1.neighbors, ep0)
	}


	step1(allNodes["start"], allNodes["end"])
	step2(allNodes["start"], allNodes["end"])
}