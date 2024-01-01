package main

import (
	"slices"

	"github.com/erdnaxeli/adventofcode/aoc"
)

func (s solver) Day23p1(input aoc.Input) string {
	grid := ToGrid(input)
	start := aoc.Point{X: 0, Y: 1}
	path := GetLongestHike(
		Path{cost: 0, points: []aoc.Point{start}},
		aoc.Point{X: len(grid) - 1, Y: len(grid[0]) - 2},
		grid,
	)
	return aoc.ResultI(path.cost)
}

func GetLongestHike(path Path, end aoc.Point, grid [][]byte) Path {
	point := path.points[len(path.points)-1]
	if point.X == end.X && point.Y == end.Y {
		return path
	}

	var nextPoints []aoc.Point
	switch grid[point.X][point.Y] {
	case '#':
		// we can't walk on trees
		return Path{}
	case '>':
		p := point
		p.Y++
		nextPoints = append(nextPoints, p)
	case '<':
		p := point
		p.Y--
		nextPoints = append(nextPoints, p)
	case '^':
		p := point
		p.X--
		nextPoints = append(nextPoints, p)
	case 'v':
		p := point
		p.X++
		nextPoints = append(nextPoints, p)
	default:
		p1, p2, p3, p4 := point, point, point, point
		p1.X--
		p2.X++
		p3.Y--
		p4.Y++
		nextPoints = append(nextPoints, p1, p2, p3, p4)
	}

	bestPath := Path{cost: path.cost}
	for _, np := range nextPoints {
		if np.X < 0 || np.Y < 0 || np.X >= len(grid) || np.Y >= len(grid[0]) {
			continue
		} else if grid[np.X][np.Y] == '#' {
			continue
		}

		if slices.Contains(path.points, np) {
			continue
		}

		nextPath := Path{
			cost:   path.cost + 1,
			points: append([]aoc.Point{}, path.points...),
		}
		nextPath.points = append(nextPath.points, np)
		resultPath := GetLongestHike(nextPath, end, grid)
		if len(resultPath.points) > 0 && resultPath.cost > bestPath.cost {
			bestPath = resultPath
		}
	}

	if len(bestPath.points) == 0 {
		bestPath.cost = 0
	}
	return bestPath
}

type Edge struct {
	Weight int
	Node   aoc.Point
}

type Graph map[aoc.Point][]Edge

func NewGraph() Graph {
	return make(Graph)
}

func (g Graph) AddEdge(n1, n2 aoc.Point, weight int) {
	g[n1] = append(g[n1], Edge{Weight: weight, Node: n2})
	g[n2] = append(g[n2], Edge{Weight: weight, Node: n1})
}

func (g Graph) RemoveEdge(n1, n2 aoc.Point) {
	for i, edge := range g[n1] {
		if edge.Node == n2 {
			g[n1] = append(g[n1][:i], g[n1][i+1:]...)
			break
		}
	}

	for i, edge := range g[n2] {
		if edge.Node == n1 {
			g[n2] = append(g[n2][:i], g[n2][i+1:]...)
			break
		}
	}
}

func (g Graph) RemoveNode(node aoc.Point) {
	// We need to copy g[node], as RemoveEdge updates g[node].
	for _, edge := range slices.Clone(g[node]) {
		g.RemoveEdge(node, edge.Node)
	}

	delete(g, node)

	// clean orphan nodes
	for n, edges := range g {
		if len(edges) == 0 {
			delete(g, n)
		}
	}
}

func (g Graph) Simplify(startNode aoc.Point) {
	// Replace chains of (n1 -w1-> n2 -w2-> n3) with (n1 -n1+n2-> n3)
	node := startNode
	queue := []aoc.Point{node}
	processed := make(map[aoc.Point]bool)

	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		checkNextNodes := true

		for checkNextNodes {
			var neighbors []Edge
			for _, edge := range g[node] {
				if !processed[edge.Node] {
					neighbors = append(neighbors, edge)
				}
			}

			checkNextNodes = false
			for _, e1 := range neighbors {
				if len(g[e1.Node]) == 1 || len(g[e1.Node]) > 2 {
					// Our neighbor as only one link, it must be the linked to the
					// current node.
					// Or, our neighbor is an intersection.
					continue
				}

				for _, e2 := range g[e1.Node] {
					if e2.Node == node {
						continue
					}

					/*
						log.Printf(
							"Replace %v -(%d)-> %v -(%d)-> %v by %v -(%d)-> %v\n",
							node, e1.Weight, e1.Node, e2.Weight, e2.Node, node, e1.Weight+e2.Weight, e2.Node,
						)
					*/
					g.AddEdge(node, e2.Node, e1.Weight+e2.Weight)
				}

				g.RemoveNode(e1.Node)
				checkNextNodes = true
			}
		}

		for _, edge := range g[node] {
			if !processed[edge.Node] {
				queue = append(queue, edge.Node)
			}
		}

		processed[node] = true
	}
}

func GridToGraph(grid [][]byte) Graph {
	graph := NewGraph()

	for x := range grid {
		for y := range grid[0] {
			if grid[x][y] == '#' {
				continue
			}

			node := aoc.Point{X: x, Y: y}

			if y < len(grid[0])-1 && grid[x][y+1] != '#' {
				graph.AddEdge(node, aoc.Point{X: x, Y: y + 1}, 1)
			}

			if x < len(grid)-1 && grid[x+1][y] != '#' {
				graph.AddEdge(node, aoc.Point{X: x + 1, Y: y}, 1)
			}
		}
	}

	return graph
}

func (s solver) Day23p2(input aoc.Input) string {
	grid := ToGrid(input)
	graph := GridToGraph(grid)
	graph.Simplify(aoc.Point{X: 0, Y: 1})

	result := GetLongestHikeGraph(
		graph,
		Path{cost: 0, points: []aoc.Point{{X: 0, Y: 1}}},
		aoc.Point{X: len(grid) - 1, Y: len(grid[0]) - 2},
	)
	return aoc.ResultI(result)
}

func GetLongestHikeGraph(g Graph, path Path, end aoc.Point) int {
	node := path.points[len(path.points)-1]
	// -1 means "no path found"
	best := -1

	for _, edge := range g[node] {
		if slices.Contains(path.points, edge.Node) {
			continue
		}

		if edge.Node == end {
			best = path.cost + edge.Weight
			continue
		}

		path := Path{
			cost:   path.cost + edge.Weight,
			points: append(path.points, edge.Node),
		}
		result := GetLongestHikeGraph(g, path, end)
		if result > best {
			best = result
		}
	}

	return best
}
