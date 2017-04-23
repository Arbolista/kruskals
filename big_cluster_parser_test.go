package kruskals

import(
  "testing"
  "fmt"
  "strconv"
)

func TestBigClusterParser1(t *testing.T) {
  binary_vertices := []string{
    "000",
    "011",
    "011",
    "111",
  }

  vertices := make([]int, len(binary_vertices))
  for vertex, bitString := range binary_vertices {
    value, _ := strconv.ParseInt(bitString, 2, 32)
    vertices[vertex] = int(value)
  }

  edges, totalVerticesGraphed := MakeEdges(1, vertices, 24)
  fmt.Println(edges)
  if len(edges) != 1 && totalVerticesGraphed != 1{
    t.Fatal("You fucked up holmes.")
  }

  edges, totalVerticesGraphed = MakeEdges(2, vertices, 24)
  fmt.Println(edges)
  if len(edges) != 2 && totalVerticesGraphed != 2 {
    t.Fatal("You fucked up holmes.")
  }

}

func TestBigClusterParser2(t *testing.T) {
  binary_vertices := []string{
    "1000000000", // 0
    "1110000000",// 1
    "1001100000",// 2
    "1000011000", // 3
    "1000000110",// 4
    "0100000001",// 5
    "0100001001",// 6
    "0100011001",// 7
    "0100111001",//8
    "0100111001",//9
  }

  vertices := make([]int, len(binary_vertices))
  for vertex, bitString := range binary_vertices {
    value, _ := strconv.ParseInt(bitString, 2, 32)
    vertices[vertex] = int(value)
  }

  edges, totalVerticesGraphed := MakeEdges(3, vertices, 24)
  fmt.Println(edges)
  if len(edges) != 9 && totalVerticesGraphed != 9{
    t.Fatal("You fucked up holmes.")
  }
}
