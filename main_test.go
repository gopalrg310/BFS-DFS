package main

import (
	"testing"
	"fmt"
	"reflect"
)


func Test_main(t *testing.T){
	edge:=new(edges)
	edge.src=0
	edge.dest=1
	edge.tree=make(map[int32][]int32)
	edge.AddEdges()
	edge.src=0
	edge.dest=2
	edge.AddEdges()
	edge.src=1
	edge.dest=2
	edge.AddEdges()
	edge.src=2
	edge.dest=0
	edge.AddEdges()
	edge.src=2
	edge.dest=3
	edge.AddEdges()
	edge.src=3
	edge.dest=3
	edge.AddEdges()

	want:=[]int32{2,0,1,3}
	output:=edge.doDFS(2)
	fmt.Println("Following is DFS from (starting from vertex 2)")
	fmt.Println(output)
	if !reflect.DeepEqual(output,want){
		t.Errorf("got %v, wanted %v", output, want)
	}
	want=[]int32{2,0,3,1}
	output=edge.doBFS(2)
	fmt.Println("Following is BFS from (starting from vertex 2)")
	fmt.Println(output)
	if !reflect.DeepEqual(output,want){
		t.Errorf("got %v, wanted %v", output, want)
	}
}