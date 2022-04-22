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

func (q *Queue) dequeueMin() *Vertex{
	if len(q.items) ==0 {
		return nil
	}
	min:= q.items[0]
	for i, v:= range q.items{
		if min.key > v.key {
			q.items = append(q.items[:i], q.items[i+1:]... )
			return v
		}
		i++
	}
	q.items = q.items[1:]
	return min
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
	length int
}

type Vertex struct {
	key int
	adjacent []*AdjacentVertexNode
}

type AdjacentVertexNode struct{
	cost int
	vertex *Vertex
}

func (g *Graph) addVertice(key int){

	if contains(g.vertices, key) {
		fmt.Println("Duplicate Vertex : ", key)
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: key})
	g.length++
}

func (g *Graph) addEdge(from, to , cost int){
	fromVertex:= g.getVertice(from)
	toVertex:= g.getVertice(to)

	if fromVertex == nil || toVertex == nil {
		fmt.Printf("\n Invalid Edge %d->%d", from, to);
	}else if containsAdjacent(fromVertex.adjacent, to) {
		fmt.Printf("\n Existing Edge %d->%d", from, to);
	}else {
		fromVertex.adjacent = append(fromVertex.adjacent, &AdjacentVertexNode{ cost: cost, vertex: toVertex})
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
func containsAdjacent(s []*AdjacentVertexNode, key int) bool {
	for _, v:= range s {
		if v.vertex.key == key {
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
			fmt.Print(" ", v.vertex.key)
		}
	}
}

func (g Graph) getIndexOfVertex(key int) int {
	for i,v:= range g.vertices  {
		if key == v.key {
			return i
		}
	}
	return -1
}

func (g *Graph) dijkstraAlgo(){
	if len(g.vertices) ==0{
		return
	}
	infinity:=9999
	q:= Queue{}
	var dist = make([]int,g.length)
	for i:=0; i< len(g.vertices); i++ {
		dist[i] = infinity
	}
	dist[0]=0
	for _, v:= range g.vertices {
			q.enqueue(v)	
	}
	for len(q.items) !=0 {
		min := q.dequeueMin()

		for _, v:= range min.adjacent {
			indexVertex:= g.getIndexOfVertex(v.vertex.key)
			indexMin:= g.getIndexOfVertex(min.key)

			fmt.Printf("\nMin: %d, Vertex: %d, cost: %d < %d  ", min.key, v.vertex.key,dist[indexMin]+ v.cost , dist[indexVertex] )
			
			if dist[indexMin]+ v.cost < dist[indexVertex]{
				dist[indexVertex] =  dist[indexMin]+ v.cost
			}
			fmt.Print(" ===> cost changed: ",dist[indexVertex])
		}
	}
	fmt.Println("\n\nCost of graph paths")
	for i:=0; i<len(g.vertices); i++ {
			if i!=0 {
			fmt.Println(g.vertices[0].key," to ",g.vertices[i].key,", Cost: ",dist[i] );
		}
	}
}


func main(){

	g := &Graph{}
	// graph.addVertice(1)
	// graph.addVertice(7)
	// graph.addVertice(2)
	// graph.addVertice(3)
	// graph.addVertice(4)
	// graph.addVertice(5)
	// graph.addVertice(6)
	// graph.addVertice(8)

	// graph.addEdge(1,4, 1)
	// graph.addEdge(1,2, 2)
	// graph.addEdge(1,3, 3)
	// graph.addEdge(2,5, 4)
	// graph.addEdge(5,6, 5)
	// graph.addEdge(4,7, 6)
	// graph.addEdge(7,8, 2)


	// graph.print()

	// graph.addVertice(1)
	// graph.addVertice(2)
	// graph.addVertice(3)
	// graph.addVertice(4)
	// graph.addVertice(5)
	// graph.addVertice(6)

	// graph.addEdge(1,2, 2)
	// graph.addEdge(1,3, 4)
	// graph.addEdge(2,4, 7)
	// graph.addEdge(3,5, 3)
	// graph.addEdge(2,3, 1)
	// graph.addEdge(5,4, 2)
	// graph.addEdge(4,6, 1)
	// graph.addEdge(5,6, 5)
	

	g.addVertice(0)
	g.addVertice(1)
	g.addVertice(2)
	g.addVertice(3)
	g.addVertice(4)
	g.addVertice(5)
	g.addVertice(6)

	g.addEdge(0, 1, 3);
	g.addEdge(0, 2, 6);
	g.addEdge(1, 0, 3);
	g.addEdge(1, 2, 2);
	g.addEdge(1, 3, 1);
	g.addEdge(2, 1, 6);
	g.addEdge(2, 1, 2);
	g.addEdge(2, 3, 1);
	g.addEdge(2, 4, 4);
	g.addEdge(2, 5, 2);
	g.addEdge(3, 1, 1);
	g.addEdge(3, 2, 1);
	g.addEdge(3, 4, 2);
	g.addEdge(3, 6, 4);
	g.addEdge(4, 2, 4);
	g.addEdge(4, 3, 2);
	g.addEdge(4, 5, 2);
	g.addEdge(4, 6, 1);
	g.addEdge(5, 2, 2);
	g.addEdge(5, 4, 2);
	g.addEdge(5, 6, 1);
	g.addEdge(6, 3, 4);
	g.addEdge(6, 4, 1);
	g.addEdge(6, 5, 1);

	g.dijkstraAlgo()

}