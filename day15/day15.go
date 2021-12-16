package day15

import (
	_ "embed"
	"fmt"
	"github.com/nickdavies/go-astar/astar"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

const INF = math.MaxInt

type Node struct {
	x int
	y int
	out bool
	value int
	distance int

	left *Node
	right *Node
	up *Node
	down *Node
}

func min(a int, b int) int {
	if a < b { return a }
	return b
}

func DrawPath(grid [][]string, path *astar.PathPoint, path_char string) {
	for {
		grid[path.Row][path.Col] = path_char

		path = path.Parent
		if path == nil {
			break
		}
	}
}

func PrintGrid(grid [][]*Node) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j].distance == INF {
				fmt.Printf(" __ ")
			} else {
				fmt.Printf("%3d ", grid[i][j].distance)
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func dijkstra(matrix [][]*Node, sortedNodes []*Node) {

	outNodes := []*Node {}
	startNode := sortedNodes[0]
	startNode.distance = startNode.value

	for i := 0; len(outNodes) < len(matrix) * len(matrix[0]); i++ {


		//PrintGrid(matrix)
		//fmt.Printf("Working on node %d, %d = %d\n", startNode.x, startNode.y, startNode.distance)

		if startNode.left != nil && ! startNode.left.out {
			startNode.left.distance = min(startNode.left.distance, startNode.distance + startNode.left.value)
		}
		if startNode.right != nil && ! startNode.right.out {
			startNode.right.distance = min(startNode.right.distance, startNode.distance + startNode.right.value)
		}
		if startNode.up != nil && ! startNode.up.out {
			startNode.up.distance = min(startNode.up.distance, startNode.distance + startNode.up.value)
		}
		if startNode.down != nil && ! startNode.down.out {
			startNode.down.distance = min(startNode.down.distance, startNode.distance + startNode.down.value)
		}

		sort.Slice(sortedNodes, func(i, j int) bool {
			return sortedNodes[i].distance < sortedNodes[j].distance
		})

		startNode.out = true
		outNodes = append(outNodes, startNode)

		startNode = sortedNodes[0]
		sortedNodes = sortedNodes[1:]
	}

	fmt.Println(matrix[0][0].distance)
	//PrintGrid(matrix)

}

func step1() {
	var matrix [][]*Node
	sortedNodes := make([]*Node, 0)

	for i, line := range strings.Split(data, "\n") {
		row := make([]*Node, len(line))
		for j, digit := range strings.Split(line, "") {
			value, _ := strconv.Atoi(digit)

			row[j] = &Node{
				x:        i,
				y:        j,
				out:      false,
				value:    value,
				distance: INF,
				left:     nil,
				right:    nil,
				up:       nil,
				down:     nil,
			}

			sortedNodes = append(sortedNodes, row[j])
		}
		matrix = append(matrix, row)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if i > 0 { matrix[i][j].up = matrix[i-1][j] }
			if j > 0 { matrix[i][j].left = matrix[i][j-1] }

			if i < len(matrix) - 1 { matrix[i][j].down = matrix[i+1][j] }
			if j < len(matrix[0]) - 1 { matrix[i][j].right = matrix[i][j+1] }

		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {


		}
	}

	matrix[0][0].value = 0
	matrix[len(matrix)-1][len(matrix[0])-1].distance = matrix[len(matrix)-1][len(matrix[0])-1].value

	sort.Slice(sortedNodes, func(i, j int) bool {
		return sortedNodes[i].distance < sortedNodes[j].distance
	})

	dijkstra(matrix, sortedNodes)
}

func cappedValue(val int) int {
	res := val
	for ; res > 9; res -= 9 {}
	return res
}

func step2() {
	matrix := make([][]*Node, 500)
	for i :=0; i < 500; i++ {
		matrix[i] = make([]*Node, 500)
	}

	sortedNodes := make([]*Node, 0)

	for i, line := range strings.Split(data, "\n") {
		for j, digit := range strings.Split(line, "") {
			value, _ := strconv.Atoi(digit)

			for ni := 0; ni < 5; ni++ {
				for nj := 0; nj < 5; nj++ {
					matrix[(100 * ni) + i][(100 * nj) + j] = &Node{
						x:        i,
						y:        j,
						out:      false,
						value:    cappedValue(value + ni + nj),
						distance: INF,
						left:     nil,
						right:    nil,
						up:       nil,
						down:     nil,
					}

					sortedNodes = append(sortedNodes, matrix[(100 * ni) + i][(100 * nj) + j])
				}
			}
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {

			if i > 0 { matrix[i][j].up = matrix[i-1][j] }
			if j > 0 { matrix[i][j].left = matrix[i][j-1] }

			if i < len(matrix) - 1 { matrix[i][j].down = matrix[i+1][j] }
			if j < len(matrix[0]) - 1 { matrix[i][j].right = matrix[i][j+1] }

		}
	}

	matrix[0][0].value = 0
	matrix[len(matrix)-1][len(matrix[0])-1].distance = matrix[len(matrix)-1][len(matrix[0])-1].value

	// Easy sort! only the last value has a distance < INF
	startNode := sortedNodes[len(sortedNodes) - 1]
	sortedNodes[len(sortedNodes) - 1] = sortedNodes[0]
	sortedNodes[0] = startNode

	dijkstra(matrix, sortedNodes)
}

func Solve() {
	step1()
	step2()
}