package main

import (
	"fmt"
	"strconv"
)


var permutation [][]int
var N = 4

func main () {

  var source = 0;
  var graph = [][]int {
		[]int{0, 10, 15, 20},
		[]int{10, 0, 35, 25},
		[]int{15, 35, 0, 30},
		[]int{20, 25, 30, 0},
	}

  tsp(graph, source)
}

func tsp(graph [][]int, source int) {

	// Set of vertices
	var vertices []int

	// best graph path - subgraph
	var vertexCover [][]int  
	var vertexCoverTemp [][]int  

	// Total path weight
	var currentPathWeight = 0;
	
	// BEST PATH TOTAL WEIGHT
	// start from 1 BLN km
	var minPathWeight = 1000000000;

  // Insert vertices of graph into array - to change
  vertices = createVertexArray(source)

  // Get all the permutation
  getPermutation(vertices, nil)

  for _, node := range permutation {

  	// Reset
  	currentPathWeight = 0
  	vertexCoverTemp = nil

    for i,n:=0,len(vertices)-1; i<n; i++ {

  		// Get arch of graph weight
  		archWeight := graph[node[i]][node[i+1]] 

      // fmt.Println("From node " + strconv.Itoa(node[i]+1) + " to node " + strconv.Itoa(node[i+1]+1) + " the cost is " + strconv.Itoa(archWeight) + " km")

      // Increment path weight
      currentPathWeight += archWeight;

      // Push node into graph path
      vertexCoverTemp = append(vertexCoverTemp, []int{ node[i], node[i+1], archWeight})
  	}

		// fmt.Println("Total weight is " + strconv.Itoa(currentPathWeight) + "\n\n")

	  /*
	  *   Find a new best path: update path-weight and path-graph
	  */

	  if (currentPathWeight < minPathWeight) {

	    minPathWeight = currentPathWeight;
      vertexCover = nil

	    for _, vertex := range vertexCoverTemp {

	    	vertexCover = append(vertexCover, vertex)
	    }
	  }

  }

  /*
   *	PRINT
   *
   */

   fmt.Println("THE BEST PATH IS LONG " + strconv.Itoa(minPathWeight) + " KM \n")

   for _, vertex := range vertexCover {

   	fmt.Println("From node " + strconv.Itoa(vertex[0]+1) + " to node " + strconv.Itoa(vertex[1]+1) + " the cost is " + strconv.Itoa(vertex[2]) + " km")
   }
}

func createVertexArray(source int) []int {

  var new []int;

  for i := 0; i < N; i++ {
    // if (i != source) 
      new = append(new, i); 
  }

  return new;
}

func getPermutation(items []int, perms []int) []int {

  if len(items) == 0 {
  
    permutation = append(permutation, perms)
    return perms;

  }  else {

  	var newPerms []int
    var newItems []int

		for i:=len(items)-1; i >= 0; i-- {
      
      newItems = items
      newPerms = perms

      foo := arraySplice(&newItems, i)
      // prepend
      newPerms = append([]int{ foo }, newPerms...)

      getPermutation(newItems, newPerms)
    }
  }

  return nil
}


func arraySplice(items *[]int, pos int) int {

  var res int;
  var newItems []int;

    for i, item := range(*items) {
    
      if i == pos {
        
        res = item

      } else {
        
        newItems = append(newItems, item)
      }

      i++
    }

    *items = newItems
    return res
} 
