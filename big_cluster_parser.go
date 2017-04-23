package kruskals

import (
  "os"
  "log"
  "strings"
  "strconv"
  "bufio"
)

func getVerticesAtHammingDistance(bitValue int, hammingDistance int, bitValueVertices map[int][]int, nBits int) []int {
  valuesAtHammingDistance := getValuesAtHammingDistance(bitValue, hammingDistance, bitValueVertices, nBits)

  var verticesAtHammingDistance []int
  for _, valueAtHammingDistance := range valuesAtHammingDistance {
    verticesAtDistance, _ := bitValueVertices[valueAtHammingDistance]
    for _, vertex := range verticesAtDistance {
      verticesAtHammingDistance = append(verticesAtHammingDistance, vertex)
    }
  }

  return verticesAtHammingDistance
}


func getValuesAtHammingDistance(bitValue int, hammingDistance int, bitValueVertices map[int][]int, nBits int) []int {
  var valuesAtHammingDistance []int
  if hammingDistance == 0 {
    return []int{bitValue}
  } else if hammingDistance == 1 {
    for i:=0; i<nBits; i++ {
      valueAtHammingDistance := FlipBit(bitValue, i)
      valuesAtHammingDistance = append(valuesAtHammingDistance, valueAtHammingDistance)
    }
  } else if hammingDistance == 2 {
    for i:=0; i<nBits; i++ {
      for j:=i+1; j<nBits; j++ {
        valueAtHammingDistance := FlipBit(bitValue, i)
        valueAtHammingDistance = FlipBit(valueAtHammingDistance, j)
        valuesAtHammingDistance = append(valuesAtHammingDistance, valueAtHammingDistance)
      }
    }
  } else if hammingDistance == 3 {
    for i:=0; i<nBits-2; i++ {
      for j:=i+1; j<nBits-1; j++ {
        for k:=j+1; k<nBits; k++ {
          valueAtHammingDistance := FlipBit(bitValue, i)
          valueAtHammingDistance = FlipBit(valueAtHammingDistance, j)
          valueAtHammingDistance = FlipBit(valueAtHammingDistance, k)
          valuesAtHammingDistance = append(valuesAtHammingDistance, valueAtHammingDistance)
        }
      }
    }
  }
  return valuesAtHammingDistance
}


func MakeEdges(maxHammingDistance int, vertices []int, nBits int) ([][]int, int) {

  bitValueVertices := make(map[int][]int)
  for vertex, bitValue := range vertices {
    bitValueVertices[bitValue] = append(bitValueVertices[bitValue], vertex)
  }

  var edges [][]int
  verticesConnectedBelowMaxDistance := make([]bool, len(vertices))
  for vertex, bitValue := range vertices { // n
    if verticesConnectedBelowMaxDistance[vertex] {
      continue
    }
    for hammingDistance := 0; hammingDistance < maxHammingDistance; hammingDistance++ { // 2
      verticesAtHammingDistance := getVerticesAtHammingDistance(bitValue, hammingDistance, bitValueVertices, nBits)
      for _, vertexAtHammingDistance := range verticesAtHammingDistance {
        if vertex == vertexAtHammingDistance {
          continue
        }
        verticesConnectedBelowMaxDistance[vertex] = true
        verticesConnectedBelowMaxDistance[vertexAtHammingDistance] = true
        edge := []int{vertex, vertexAtHammingDistance, hammingDistance}
        edges = append(edges, edge)
      }
    }
  }
  verticesGraphed := 0
  for i:=0; i<len(verticesConnectedBelowMaxDistance); i++{
    if verticesConnectedBelowMaxDistance[i] {
      verticesGraphed += 1
    }
  }

  return edges, verticesGraphed
}

func ParseBigEdges(filename string, maxHammingDistance int) (int, int, [][]int) {
  file, err := os.Open(filename)
  defer file.Close()
  totalVertices := 0
  nBits := 0
  if err != nil {
      log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)
  var vertices []int
  for scanner.Scan() {
    s := scanner.Text()
    parts := strings.Split(s, " ")
    if len(parts) == 2 {
      totalVertices, _ = strconv.Atoi(parts[0])
      nBits, _ = strconv.Atoi(parts[1])
      continue
    } else if len(parts) > 2 {
      binary := strings.Join(parts, "")
      value, _ := strconv.ParseInt(binary, 2, 32)
      vertices = append(vertices, int(value))
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  edges, verticesGraphed := MakeEdges(maxHammingDistance, vertices, nBits)
  return totalVertices, verticesGraphed, edges
}
