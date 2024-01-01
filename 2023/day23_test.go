package main

import (
	"testing"

	"github.com/erdnaxeli/adventofcode/aoc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDay23p1(t *testing.T) {
	input := aoc.NewInput(`#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`)

	result := solver{}.Day23p1(input)

	assert.Equal(t, "94", result)
}

func TestDay23p2(t *testing.T) {
	input := aoc.NewInput(`#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`)

	result := solver{}.Day23p2(input)

	assert.Equal(t, "154", result)
}

func TestGridToGraph(t *testing.T) {
	grid := [][]byte{
		{'#', '.', '.', '.', '#'},
		{'#', '.', '#', '#', '#'},
		{'#', '.', '.', '.', '.'},
	}

	graph := GridToGraph(grid)

	expected := NewGraph()
	expected.AddEdge(aoc.Point{X: 0, Y: 1}, aoc.Point{X: 0, Y: 2}, 1)
	expected.AddEdge(aoc.Point{X: 0, Y: 2}, aoc.Point{X: 0, Y: 3}, 1)
	expected.AddEdge(aoc.Point{X: 0, Y: 1}, aoc.Point{X: 1, Y: 1}, 1)
	expected.AddEdge(aoc.Point{X: 1, Y: 1}, aoc.Point{X: 2, Y: 1}, 1)
	expected.AddEdge(aoc.Point{X: 2, Y: 1}, aoc.Point{X: 2, Y: 2}, 1)
	expected.AddEdge(aoc.Point{X: 2, Y: 2}, aoc.Point{X: 2, Y: 3}, 1)
	expected.AddEdge(aoc.Point{X: 2, Y: 3}, aoc.Point{X: 2, Y: 4}, 1)
	assert.Equal(t, expected, graph)
}

func TestGraphRemoveNode(t *testing.T) {
	a := aoc.Point{X: 0, Y: 0}
	b := aoc.Point{X: 1, Y: 0}
	c := aoc.Point{X: 2, Y: 0}
	d := aoc.Point{X: 3, Y: 0}
	graph := NewGraph()
	graph.AddEdge(a, b, 1)
	graph.AddEdge(b, c, 1)
	graph.AddEdge(b, d, 1)
	graph.AddEdge(c, d, 1)

	graph.RemoveNode(b)

	expected := NewGraph()
	expected.AddEdge(c, d, 1)
	assert.Equal(t, expected, graph)
}

func TestGrapSimplify_simple(t *testing.T) {
	a := aoc.Point{X: 0, Y: 0}
	b := aoc.Point{X: 1, Y: 0}
	c := aoc.Point{X: 2, Y: 0}
	d := aoc.Point{X: 3, Y: 0}
	graph := NewGraph()
	graph.AddEdge(a, b, 1)
	graph.AddEdge(b, c, 1)
	graph.AddEdge(c, d, 1)

	graph.Simplify(a)

	expected := NewGraph()
	expected.AddEdge(a, d, 3)
	assert.Equal(t, expected, graph)
}

func TestGraphSimplify_twoBranches(t *testing.T) {
	a := aoc.Point{X: 1, Y: 0}
	b := aoc.Point{X: 2, Y: 0}
	c := aoc.Point{X: 3, Y: 0}
	d := aoc.Point{X: 4, Y: 0}
	e := aoc.Point{X: 0, Y: 1}
	f := aoc.Point{X: 0, Y: 2}
	g := aoc.Point{X: 0, Y: 3}
	// a -> b -> c -> d
	//  \
	//   -> e -> f -> g
	graph := NewGraph()
	graph.AddEdge(a, b, 1)
	graph.AddEdge(b, c, 1)
	graph.AddEdge(c, d, 1)
	graph.AddEdge(a, e, 1)
	graph.AddEdge(e, f, 1)
	graph.AddEdge(f, g, 1)

	graph.Simplify(a)

	// a -3-> d
	//  \
	//   -3-> g
	expected := NewGraph()
	expected.AddEdge(a, d, 3)
	expected.AddEdge(a, g, 3)
	assert.Equal(t, expected, graph)
}

func TestGraphSimplify_moreBranches(t *testing.T) {
	a := aoc.Point{X: 1}
	b := aoc.Point{X: 2}
	c := aoc.Point{X: 3}
	d := aoc.Point{X: 4}
	e := aoc.Point{X: 5}
	f := aoc.Point{X: 6}
	g := aoc.Point{X: 7}
	h := aoc.Point{X: 8}
	i := aoc.Point{X: 9}
	j := aoc.Point{X: 10}
	k := aoc.Point{X: 11}
	l := aoc.Point{X: 12}
	m := aoc.Point{X: 13}
	n := aoc.Point{X: 14}
	o := aoc.Point{X: 15}
	p := aoc.Point{X: 16}
	q := aoc.Point{X: 17}
	// a -> b -> c -> d -> e -> f -> l
	//            \-> o -> p -> q -> g -> h -> i -> m
	//                           \-> j -> k -> n
	graph := NewGraph()
	graph.AddEdge(a, b, 1)
	graph.AddEdge(b, c, 1)
	graph.AddEdge(c, d, 1)
	graph.AddEdge(d, e, 1)
	graph.AddEdge(e, f, 1)
	graph.AddEdge(f, l, 1)
	graph.AddEdge(c, o, 1)
	graph.AddEdge(o, p, 1)
	graph.AddEdge(p, q, 1)
	graph.AddEdge(q, g, 1)
	graph.AddEdge(g, h, 1)
	graph.AddEdge(h, i, 1)
	graph.AddEdge(i, m, 1)
	graph.AddEdge(q, j, 1)
	graph.AddEdge(j, k, 1)
	graph.AddEdge(k, n, 1)

	graph.Simplify(a)

	// a -2-> c -4-> l
	//         \-3-> q -4-> m
	//                \-3-> n
	expected := NewGraph()
	expected.AddEdge(a, c, 2)
	expected.AddEdge(c, l, 4)
	expected.AddEdge(c, q, 3)
	expected.AddEdge(q, m, 4)
	expected.AddEdge(q, n, 3)

	// assert.Equal(t, expected, graph)
	require.Len(t, graph, len(expected))
	for node, edges := range expected {
		gEdges, ok := graph[node]
		require.True(t, ok)
		assert.ElementsMatch(t, edges, gEdges)
	}
}

func TestGraphSimplify_example(t *testing.T) {
	input := aoc.NewInput(`#.#####################
#.......#########...###
#######.#########.#.###
###.....#.>.>.###.#.###
###v#####.#v#.###.#.###
###.>...#.#.#.....#...#
###v###.#.#.#########.#
###...#.#.#.......#...#
#####.#.#.#######.#.###
#.....#.#.#.......#...#
#.#####.#.#.#########v#
#.#...#...#...###...>.#
#.#.#v#######v###.###v#
#...#.>.#...>.>.#.###.#
#####v#.#.###v#.#.###.#
#.....#...#...#.#.#...#
#.#########.###.#.#.###
#...###...#...#...#.###
###.###.#.###v#####v###
#...#...#.#.>.>.#.>.###
#.###.###.#.###.#.#v###
#.....###...###...#...#
#####################.#`)
	grid := ToGrid(input)
	graph := GridToGraph(grid)

	graph.Simplify(aoc.Point{X: 0, Y: 1})

	/*
		#0#####################
		#.......#########...###
		#######.#########.#.###
		###.....#..O..###.#.###
		###.#####.#.#.###.#.###
		###0....#.#.#.....#...#
		###.###.#.#.#########.#
		###...#.#.#.......#...#
		#####.#.#.#######.#.###
		#.....#.#.#.......#...#
		#.#####.#.#.#########.#
		#.#...#...#...###....O#
		#.#.#.#######.###.###.#
		#...#O..#....O..#.###.#
		#####.#.#.###.#.#.###.#
		#.....#...#...#.#.#...#
		#.#########.###.#.#.###
		#...###...#...#...#.###
		###.###.#.###.#####.###
		#...#...#.#..O..#..O###
		#.###.###.#.###.#.#.###
		#.....###...###...#...#
		#####################0#
	*/
	expected := NewGraph()
	expected.AddEdge(aoc.Point{X: 0, Y: 1}, aoc.Point{X: 5, Y: 3}, 15)
	expected.AddEdge(aoc.Point{X: 5, Y: 3}, aoc.Point{X: 13, Y: 5}, 22)
	expected.AddEdge(aoc.Point{X: 5, Y: 3}, aoc.Point{X: 3, Y: 11}, 22)
	expected.AddEdge(aoc.Point{X: 13, Y: 5}, aoc.Point{X: 19, Y: 13}, 38)
	expected.AddEdge(aoc.Point{X: 13, Y: 5}, aoc.Point{X: 13, Y: 13}, 12)
	expected.AddEdge(aoc.Point{X: 3, Y: 11}, aoc.Point{X: 13, Y: 13}, 24)
	expected.AddEdge(aoc.Point{X: 3, Y: 11}, aoc.Point{X: 11, Y: 21}, 30)
	expected.AddEdge(aoc.Point{X: 19, Y: 13}, aoc.Point{X: 19, Y: 19}, 10)
	expected.AddEdge(aoc.Point{X: 19, Y: 13}, aoc.Point{X: 13, Y: 13}, 10)
	expected.AddEdge(aoc.Point{X: 13, Y: 13}, aoc.Point{X: 11, Y: 21}, 18)
	expected.AddEdge(aoc.Point{X: 11, Y: 21}, aoc.Point{X: 19, Y: 19}, 10)
	expected.AddEdge(aoc.Point{X: 19, Y: 19}, aoc.Point{X: 22, Y: 21}, 5)

	require.Len(t, graph, len(expected))
	for node, edges := range expected {
		gEdges, ok := graph[node]
		require.True(t, ok)
		assert.ElementsMatch(t, edges, gEdges)
	}
}
