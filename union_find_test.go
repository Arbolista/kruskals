package kruskals

import(
  "testing"
  "fmt"
  "log"
)

func assertUfState(t *testing.T, ufd *UnionFind, expectedRanks []int, expectedParents []int, expectedClusters int){
  if ufd.Count() != expectedClusters {
    msg := fmt.Sprintf("Expected %d clusters, got %d", expectedClusters, ufd.Count())
    panic(msg)
  }
  for i, expectedParent := range expectedParents {
    realParent := ufd.Find(i)
    if realParent != expectedParent {
      msg := fmt.Sprintf("Expected parent of vertex %d to be %d, got %d", i, expectedParent, realParent)
      panic(msg)
    }
  }
  for i, expectedRank := range expectedRanks {
    realRank := ufd.Ranks[i]
    if realRank != expectedRank {
      msg := fmt.Sprintf("Expected rank of vertex %d to be %d, got %d", i, expectedRank, realRank)
      panic(msg)
    }
  }
}

func TestUnionFind1(t *testing.T) {

  log.Println(fmt.Sprintf("poo gas"))

  ufd := NewUnionFind(9)
  expectedRanks := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
  expectedParents := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
  assertUfState(t, ufd, expectedRanks, expectedParents, 9)

  ufd.Union(0, 1)
  expectedRanks = []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 2, 3, 4, 5, 6, 7, 8}
  log.Println(fmt.Sprintf("Count() %d", ufd.Count()))
  assertUfState(t, ufd, expectedRanks, expectedParents, 8)

  ufd.Union(1, 0)
  expectedRanks = []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 2, 3, 4, 5, 6, 7, 8}
  assertUfState(t, ufd, expectedRanks, expectedParents, 8)

  ufd.Union(0, 2)
  expectedRanks = []int{1, 0, 0, 0, 0, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 4, 5, 6, 7, 8}
  assertUfState(t, ufd, expectedRanks, expectedParents, 7)

  ufd.Union(4, 7)
  expectedRanks = []int{1, 0, 0, 0, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 4, 5, 6, 4, 8}
  assertUfState(t, ufd, expectedRanks, expectedParents, 6)

  ufd.Union(4, 8)
  expectedRanks = []int{1, 0, 0, 0, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 4, 5, 6, 4, 4}
  assertUfState(t, ufd, expectedRanks, expectedParents, 5)

  ufd.Union(3, 5)
  expectedRanks = []int{1, 0, 0, 1, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 4, 3, 6, 4, 4}
  assertUfState(t, ufd, expectedRanks, expectedParents, 4)

  ufd.Union(3, 4)
  expectedRanks = []int{1, 0, 0, 2, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 3, 3, 6, 3, 3}
  assertUfState(t, ufd, expectedRanks, expectedParents, 3)

  ufd.Union(5, 4)
  expectedRanks = []int{1, 0, 0, 2, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 3, 3, 6, 3, 3}
  assertUfState(t, ufd, expectedRanks, expectedParents, 3)

  ufd.Union(3, 6)
  expectedRanks = []int{1, 0, 0, 2, 1, 0, 0, 0, 0}
  expectedParents = []int{0, 0, 0, 3, 3, 3, 3, 3, 3}
  assertUfState(t, ufd, expectedRanks, expectedParents, 2)

  ufd.Union(0, 3)
  expectedRanks = []int{1, 0, 0, 2, 1, 0, 0, 0, 0}
  expectedParents = []int{3, 3, 3, 3, 3, 3, 3, 3, 3}
  assertUfState(t, ufd, expectedRanks, expectedParents, 1)
}
