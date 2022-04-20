package main

import "fmt"


type Queue struct {
	items []*Vertex
}

func (q *Queue) enqueue(v *Vertex){
	q.items = append(q.items, v)
}

func (q *Queue) dequeue() *Vertex{
	if len(q.items) ==0 {
		return nil
	}
	v:= q.items[0]
	q.items = q.items[1:]
	return v

}

type Stack struct {
	items []*Vertex
}

func (s *Stack) push(v *Vertex){
	s.items = append(s.items, v)
}

func (s *Stack) pop() *Vertex{
	if len(s.items) ==0 {
		return nil
	}
	l:= len(s.items)-1
	v:= s.items[l]
	s.items = s.items[:l]
	return v
}



type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key int
	adjacent []*Vertex
}

func (g *Graph) addVertice(key int){

	if contains(g.vertices, key) {
		fmt.Println("Duplicate Vertex : ", key)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
}

func (g *Graph) addEdge(from, to int){
	fromVertex:= g.getVertice(from)
	toVertex:= g.getVertice(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("\n Invalid Edge %d->%d", from, to);
	}else if contains(fromVertex.adjacent, to) {
		fmt.Printf("\n Existing Edge %d->%d", from, to);
	}else {
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}

}

func contains(s []*Vertex, key int) bool {
	for _, v:= range s {
		if v.key == key {
			return true
		}
	}
	return false
}

func (g *Graph) getVertice(key int) *Vertex {
	for i, v:= range g.vertices {
		if v.key == key {
			return g.vertices[i]
		}
	}
	return nil
}

func (g *Graph) print(){
	for _,v := range g.vertices {
		fmt.Print("\n", v.key, ":")
		for _,v:= range v.adjacent {
			fmt.Print(" ", v.key)
		}
	}
}

func (g *Graph) bfs(){
	q:= Queue{}
	visited:= make([]int,0)
	q.enqueue(g.vertices[0])
	for len(q.items) >0 {
		n:= q.dequeue()
		visited = append(visited, n.key)
		for _, v:= range n.adjacent {
			q.enqueue(v)
		}
	}
	fmt.Println("BFS: ", visited)
}

func (g *Graph) dfs(){
	s:= Stack{}
	visited:= make([]int,0)
	s.push(g.vertices[0])
	for len(s.items) >0 {
		n:= s.pop()
		visited = append(visited, n.key)
	
		for _, v:= range n.adjacent {
			s.push(v)
		}
	}
	fmt.Println("DFS: ", visited)
}



func main(){
	graph := &Graph{}
	graph.addVertice(1)
	graph.addVertice(2)
	graph.addVertice(3)
	graph.addVertice(4)
	graph.addVertice(5)
	graph.addVertice(6)
	graph.addVertice(7)
	graph.addVertice(7)

	graph.addEdge(1,4)
	graph.addEdge(1,2)
	graph.addEdge(1,3)
	graph.addEdge(2,5)
	graph.addEdge(5,6)
	graph.addEdge(4,7)
	graph.print()
	graph.bfs()
	graph.dfs()

}