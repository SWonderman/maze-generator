package main

import (
	"math/rand"
)

type Edge struct {
	From   *GridNode
	To     *GridNode
	Weight float64
}

type GridNode struct {
	Row    int
	Column int
}

func (node *GridNode) GetWallNeighbours(matrix *[][]byte, wallChar byte) []*Edge {
	var edges []*Edge

	row := node.Row
	column := node.Column

	orthogonalDirectionWeight := 1.0

	// TOP
	if row-1 >= 0 && (*matrix)[row-1][column] == wallChar {
		edges = append(edges, &Edge{node, &GridNode{Row: row - 1, Column: column}, orthogonalDirectionWeight})
	}

	// DOWN
	if row+1 <= len(*matrix)-1 && (*matrix)[row+1][column] == wallChar {
		edges = append(edges, &Edge{node, &GridNode{Row: row + 1, Column: column}, orthogonalDirectionWeight})
	}

	// RIGHT
	if column+1 <= len((*matrix)[0])-1 && (*matrix)[row][column+1] == wallChar {
		edges = append(edges, &Edge{node, &GridNode{Row: row, Column: column + 1}, orthogonalDirectionWeight})
	}

	// LEFT
	if column-1 >= 0 && (*matrix)[row][column-1] == wallChar {
		edges = append(edges, &Edge{node, &GridNode{Row: row, Column: column - 1}, orthogonalDirectionWeight})
	}

	return edges
}

func (node *GridNode) GetNeighbours(matrix *[][]byte, wallChar byte) []*Edge {
	var edges []*Edge

	row := node.Row
	column := node.Column

	orthogonalDirectionWeight := 1.0

	// TOP
	if row-1 >= 0 {
		edges = append(edges, &Edge{node, &GridNode{Row: row - 1, Column: column}, orthogonalDirectionWeight})
	}

	// DOWN
	if row+1 <= len(*matrix)-1 {
		edges = append(edges, &Edge{node, &GridNode{Row: row + 1, Column: column}, orthogonalDirectionWeight})
	}

	// RIGHT
	if column+1 <= len((*matrix)[0])-1 {
		edges = append(edges, &Edge{node, &GridNode{Row: row, Column: column + 1}, orthogonalDirectionWeight})
	}

	// LEFT
	if column-1 >= 0 {
		edges = append(edges, &Edge{node, &GridNode{Row: row, Column: column - 1}, orthogonalDirectionWeight})
	}

	return edges
}

type PrimsNode struct {
	CurrentNode *GridNode
	LastEdge    *Edge
	Backtrack   *PrimsNode
	CostToNode  float64
}

func (node *PrimsNode) GetCost() float64 {
	return node.CostToNode
}

func containsNode(haystack []*GridNode, needle *GridNode) bool {
	contains := false
	for _, n := range haystack {
		if n.Row == needle.Row && n.Column == needle.Column {
			contains = true
		}
	}
	return contains
}

// 'Default' implementation of Prim's
func RunPrims(matrix *[][]byte, start *GridNode, obstacleChar byte) []*Edge {
	visited := []*GridNode{}

	pqueue := &MinHeap{}
	startPrimsNode := &PrimsNode{start, nil, nil, 0}

	pqueue.Push(startPrimsNode)

	minimumSpanningTree := []*Edge{}

	for _, edge := range start.GetNeighbours(matrix, obstacleChar) {
		pqueue.Push(&PrimsNode{edge.To, edge, startPrimsNode, edge.Weight})
	}

	for pqueue.Len() > 0 {
		currentPrimsNode := pqueue.Pop().(*PrimsNode)

		// Skip when the current node was already visited
		if containsNode(visited, currentPrimsNode.CurrentNode) {
			continue
		}
		visited = append(visited, currentPrimsNode.CurrentNode)

		if currentPrimsNode.LastEdge != nil {
			minimumSpanningTree = append(minimumSpanningTree, currentPrimsNode.LastEdge)
		}

		for _, edge := range currentPrimsNode.CurrentNode.GetNeighbours(matrix, obstacleChar) {
			if !containsNode(visited, edge.To) {
				pqueue.Push(&PrimsNode{edge.To, edge, currentPrimsNode, edge.Weight})
			}
		}
	}

	return minimumSpanningTree
}

func RunMazeGeneratingPrims(matrix *[][]byte, start *GridNode, obstacleChar byte) []*GridNode {
	visited := []*GridNode{}
    passages := []*GridNode{}
	// NOTE: In this version of the algorithm, the pqueue is really just a list/array
	// from which values will be drawn at 'random'
	frontiers := []*PrimsNode{}

	startPrimsNode := &PrimsNode{start, nil, nil, 0}
	(*matrix)[startPrimsNode.CurrentNode.Row][startPrimsNode.CurrentNode.Column] = '+'

	frontiers = append(frontiers, startPrimsNode)

	for _, edge := range start.GetNeighbours(matrix, obstacleChar) {
		primsNode := &PrimsNode{edge.To, edge, startPrimsNode, edge.Weight}
		frontiers = append(frontiers, primsNode)
	}

	for len(frontiers) > 0 {
		randomFrontierIdx := rand.Intn(len(frontiers))
		frontier := frontiers[randomFrontierIdx]
		frontiers = append(frontiers[:randomFrontierIdx], frontiers[randomFrontierIdx+1:]...)

		passagesAroundTheFrontier := 0
		for _, edge := range frontier.CurrentNode.GetNeighbours(matrix, obstacleChar) {
			if (*matrix)[edge.To.Row][edge.To.Column] == '+' {
				passagesAroundTheFrontier += 1
			}
		}

		if containsNode(visited, frontier.CurrentNode) {
			continue
		}

		if passagesAroundTheFrontier == 1 {
			visited = append(visited, frontier.CurrentNode)

			(*matrix)[frontier.CurrentNode.Row][frontier.CurrentNode.Column] = '+'
            passages = append(passages, frontier.CurrentNode)

			for _, edge := range frontier.CurrentNode.GetNeighbours(matrix, obstacleChar) {
				if !containsNode(visited, edge.To) {
					primsNode := &PrimsNode{edge.To, edge, frontier, edge.Weight}
					frontiers = append(frontiers, primsNode)
				}
			}
		}
	}

    return passages
}
