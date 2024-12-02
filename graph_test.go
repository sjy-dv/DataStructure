package datastructure_test

import (
	"fmt"
	"testing"
)

/*
    Graph Structure
       A
    B    C
     <- means vertex [ A, B, C]

        A
      /
    B    C
     <- means
                vertex  A  B  C
                edge    B  A

      A         <- means
    /  \        vertex   A      B       C
   B - C        edge  [B, C]  [A, C]  [A, B]
*/

type Graph struct {
	vertexes map[any][]any
}

func NewGraph() *Graph {
	return &Graph{vertexes: make(map[any][]any)}
}

func (graph *Graph) addVertex(vertex any) bool {
	// check already exists vertex
	// vertex key not duplicate
	if _, ok := graph.vertexes[vertex]; !ok {
		//if not, initial vertex
		graph.vertexes[vertex] = make([]any, 0)
		return true
	}
	// if already exists vertext return false
	return false
}

func (graph *Graph) addEdge(vertexA, vertexB any) bool {
	_, okA := graph.vertexes[vertexA]
	_, okB := graph.vertexes[vertexB]
	// add Edge Condition is two vertex is already exists
	if okA && okB {
		// vertex A edge insert vertexB
		graph.vertexes[vertexA] = append(graph.vertexes[vertexA], vertexB)
		graph.vertexes[vertexB] = append(graph.vertexes[vertexB], vertexA)
		// now graph change view
		/*
		   Before                      A
		   Vertex  A   B
		   Edge    []  []            B

		   After                      A
		   Vertex A   B              /     <- connecting
		   Edge  [B] [A]            B
		*/
		return true
	}
	return false
}

func (graph *Graph) removeEdge(vertexA, vertexB any) bool {
	_, okA := graph.vertexes[vertexA]
	_, okB := graph.vertexes[vertexB]
	if okA && okB {
		// you can rewrite more best code
		// using type Vertex any <-
		// func(v Vertex) remove() try!!
		graph.vertexes[vertexA] = remove(graph.vertexes[vertexA], vertexB)
		graph.vertexes[vertexB] = remove(graph.vertexes[vertexB], vertexA)
		return true
	}
	return false
}

func (graph *Graph) removeVertex(vertex any) bool {
	_, ok := graph.vertexes[vertex]
	if ok {
		for _, edgeVertex := range graph.vertexes[vertex] {
			// remove link in target vertex edge vertex
			// if A vertex have edges [B, C]
			// B,C maybe has [A,...]
			// iter B -> C Edge remove A
			graph.vertexes[edgeVertex] = remove(graph.vertexes[edgeVertex], vertex)
		}
		// last step, delete vertex in graph
		delete(graph.vertexes, vertex)
		return true
	}
	return false
}

func remove(list []any, x any) []any {
	newList := make([]any, 0, len(list)-1)
	for _, ent := range list {
		if ent == x {
			continue
		}
		newList = append(newList, ent)
	}
	return newList
}

func (graph *Graph) printGraph() {
	for vertex := range graph.vertexes {
		fmt.Println(vertex, " : ", graph.vertexes[vertex])
	}
}

func TestGraph(t *testing.T) {
	graph := NewGraph()

	graph.addVertex("A")
	graph.addVertex(101)
	graph.addVertex("V")
	graph.addVertex("X")
	graph.addVertex("CB0")

	graph.addEdge("A", 101)
	graph.addEdge("A", "CB0")
	graph.addEdge("V", "A")
	graph.addEdge("X", "CB0")
	graph.addEdge(101, "V")
	graph.addEdge("X", 101)

	graph.printGraph()

	// CB0  :  [A X]
	// A  :  [101 CB0 V]
	// 101  :  [A V X]
	// V  :  [A 101]
	// X  :  [CB0 101]
	fmt.Println("========================================")
	// edge elements  3 to 2
	graph.removeEdge("A", "V")
	graph.removeEdge(101, "X")
	graph.removeEdge(101, 238192) // test not exists vertex (expect: not changed)
	graph.printGraph()
	// A  :  [101 CB0]
	// 101  :  [A V]
	// V  :  [101]
	// X  :  [CB0]
	// CB0  :  [A X]
	fmt.Println("=======================")
	// remove A vertex
	graph.removeVertex("A")
	graph.printGraph()
	// CB0  :  [X]
	// 101  :  [V]
	// V  :  [101]
	// X  :  [CB0]
}
