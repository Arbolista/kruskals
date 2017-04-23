package kruskals

import (
  //"fmt"
  //"log"
)


func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

// UnionFind is an implementation of the said data structure.
type UnionFind struct {
  Parents   []int
  Ranks   []int
  count int
}

// NewUnionFind initializes a UnionFind.
func NewUnionFind(count int) *UnionFind {
  ufd := UnionFind{count: count}
  ufd.Parents = makeRange(0, count)
  ufd.Ranks = make([]int, count)
  return &ufd
}

// Find returns the parent of the passed vertex.
// Also caches the parent on the queried vertex to
// implement path compression.
func (ufd *UnionFind) Find(vertex int) int {
  if (ufd.Parents[vertex] != vertex){
      ufd.Parents[vertex] = ufd.Find(ufd.Parents[vertex])
  }
  return ufd.Parents[vertex];
}

func (ufd *UnionFind) incrementcount() {
  ufd.count--
}

func (ufd *UnionFind) Count() int {
   return ufd.count
}

func (ufd *UnionFind) RemainingClusters() int {
  countedClusters := make([]bool, len(ufd.Parents))
  totalClusters := 0
  for i:=0; i<len(ufd.Parents); i++ {
    if !countedClusters[ufd.Find(i)] {
      countedClusters[ufd.Find(i)] = true
      totalClusters += 1
    }
  }
  return totalClusters
}

// Union adds a new vertex to the graph structure.
func (ufd *UnionFind) Union(v1 int, v2 int) {
  parent1 := ufd.Find(v1);
  parent2 := ufd.Find(v2);
  if (parent1 == parent2){ return;}

  if (ufd.Ranks[parent1] < ufd.Ranks[parent2]) {
    ufd.Parents[parent1] = parent2
    //ufd.Ranks[parent2] = ufd.Ranks[parent2] + ufd.Ranks[parent1];
  } else {
    ufd.Parents[parent2] = parent1
    if ufd.Ranks[parent1] == ufd.Ranks[parent2] {
      ufd.Ranks[parent1] = ufd.Ranks[parent2] + 1;
    }
    //ufd.Ranks[parent1] = ufd.Ranks[parent1] + ufd.Ranks[parent2];
  }
  ufd.count = ufd.count - 1;
}
