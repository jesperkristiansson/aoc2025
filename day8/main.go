package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

type point struct {
	x, y, z int
}

func getPoints() []point {
	var points []point

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		var point point
		n, err := fmt.Sscanf(line, "%d,%d,%d", &point.x, &point.y, &point.z)
		if err != nil {
			panic(err)
		}
		if n != 3 {
			panic(fmt.Errorf("failure during scan(), expected 3 items, got %d", n))
		}

		points = append(points, point)
	}

	return points
}

type distance struct {
	from, to int
	dist     float64
}

func distanceBetween(p1, p2 point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	dz := p2.z - p1.z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

func findAllDistances(points []point) []distance {
	n := len(points)
	numPairs := n * (n - 1) / 2
	distances := make([]distance, 0, numPairs)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dist := distanceBetween(points[i], points[j])
			d := distance{from: i, to: j, dist: dist}
			distances = append(distances, d)
		}
	}

	sortFunc := func(d1, d2 distance) int {
		return cmp.Compare(d1.dist, d2.dist)
	}
	slices.SortFunc(distances, sortFunc)

	return distances
}

type unionFindNode struct {
	parent int
	size   int
}

func createUnionFindNodes(n int) []unionFindNode {
	nodes := make([]unionFindNode, n)
	for i := range nodes {
		nodes[i].parent = i
		nodes[i].size = 1
	}

	return nodes
}

func findRoot(nodes []unionFindNode, i int) int {
	if nodes[i].parent != i {
		nodes[i].parent = findRoot(nodes, nodes[i].parent)
	}

	return nodes[i].parent
}

func merge(nodes []unionFindNode, i, j int) int {
	rootI := findRoot(nodes, i)
	rootJ := findRoot(nodes, j)

	if rootI == rootJ {
		return rootI
	}

	if nodes[rootI].size > nodes[rootJ].size {
		nodes[rootJ].parent = rootI
		nodes[rootI].size += nodes[rootJ].size
		nodes[rootJ].size = 0
		return rootI
	} else {
		nodes[rootI].parent = rootJ
		nodes[rootJ].size += nodes[rootI].size
		nodes[rootI].size = 0
		return rootJ
	}
}

func part1(points []point, pairsToJoin int) int {
	distances := findAllDistances(points)
	unionFindNodes := createUnionFindNodes(len(points))

	for _, d := range distances[:pairsToJoin] {
		p1 := d.from
		p2 := d.to

		merge(unionFindNodes, p1, p2)
	}

	slices.SortFunc(unionFindNodes, func(a, b unionFindNode) int { return b.size - a.size })

	return unionFindNodes[0].size * unionFindNodes[1].size * unionFindNodes[2].size
}

func part2(points []point) int {
	distances := findAllDistances(points)
	unionFindNodes := createUnionFindNodes(len(points))

	for _, d := range distances {
		p1 := d.from
		p2 := d.to

		newRoot := merge(unionFindNodes, p1, p2)
		if unionFindNodes[newRoot].size == len(unionFindNodes) {
			x1 := points[p1].x
			x2 := points[p2].x
			return x1 * x2
		}
	}

	return -1
}

func getNumPairsToJoin() int {
	if len(os.Args) < 2 {
		return 10
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("'%s' is not a number", os.Args[1]))
	}
	return n
}

func main() {
	pairsToJoin := getNumPairsToJoin()
	points := getPoints()
	fmt.Println("part 1:", part1(points, pairsToJoin))
	fmt.Println("part 2:", part2(points))
}
