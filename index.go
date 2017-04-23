package kruskals

import (
  "sort"
  "fmt"
)

// EdgeSorter sorts an array of edges with their costs in ascending order.
type edgeSorter [][]int
func (a edgeSorter) Len() int           { return len(a) }
func (a edgeSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a edgeSorter) Less(i, j int) bool {
  return a[i][2] < a[j][2]
}

// MaxDistanceKClusters returns the max spacing for k clusters of the
// given array of edges.
func MaxSpacingKClusters(k int, nVertices int, edges [][]int) int {
  sort.Sort(edgeSorter(edges))
  uf := NewUnionFind(nVertices)
  maxSpacing := 0
  for i, edge := range edges {
    if uf.Find(edge[0]) != uf.Find(edge[1]){
      uf.Union(edge[0], edge[1])
      if uf.Count() < k {
        maxSpacing = edges[i][2]
        break;
      }

    }
  }
  return maxSpacing
}

func MaxClustersNSpacing(maxSpace int, nVertices int, edges [][]int) int {
  sort.Sort(edgeSorter(edges))
  uf := NewUnionFind(nVertices)
  fmt.Println(len(uf.Parents))
  for _, edge := range edges {
    if uf.Find(edge[0]) != uf.Find(edge[1]){
      if edge[2] == maxSpace {
        break
      } else {
        uf.Union(edge[0], edge[1])
      }
    }
  }
  return uf.Count()
}
