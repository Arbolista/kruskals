package kruskals

import(
  "testing"
  "fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
  "log"
)


func TestMaxSpacingKCluster0(t *testing.T) {

  edges := [][]int{
    []int{3, 2, 60},
    []int{7, 8, 10},
    []int{6, 7, 30},
    []int{7, 5, 65},
    []int{3, 4, 80},
    []int{2, 1, 40},
    []int{6, 5, 50},
    []int{4, 8, 130},
    []int{9, 3, 5},
    []int{0, 2, 55},
    []int{4, 5, 70},
    []int{0, 1, 20},
    []int{3, 0, 100},
  }

  maxSpacing := MaxSpacingKClusters(2, 10, edges)
  if maxSpacing != 80 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(3, 10, edges)
  if maxSpacing != 70 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=3 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(4, 10, edges)
  if maxSpacing != 60 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=4 %d", maxSpacing))
  }
}


func parseEdges(filename string) (int, [][]int){

  file, err := os.Open(filename)
  defer file.Close()
  totalVertices := 0
  if err != nil {
      log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)
  var edges [][]int
  for scanner.Scan() {
    s := scanner.Text()
    parts := strings.Split(s, " ")
    if len(parts) == 1 {
      totalVertices, _ = strconv.Atoi(parts[0])
      continue
    } else if len(parts) > 1 {
      vertex1, _ := strconv.Atoi(parts[0])
      vertex2, _ := strconv.Atoi(parts[1])
      cost, _ := strconv.Atoi(parts[2])
      edge := []int{vertex1-1, vertex2-1, cost}
      edges = append(edges, edge)
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return totalVertices, edges
}

func TestMaxSpacingKClusters(t *testing.T) {

  nVertices, edges := parseEdges("./clustering.txt")
  maxSpacing := MaxSpacingKClusters(4, nVertices, edges)
  if maxSpacing != 106 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }
}

func TestMaxSpacingKClusters1(t *testing.T) {

  nVertices, edges := parseEdges("./clustering1.txt")
  maxSpacing := MaxSpacingKClusters(2, nVertices, edges)
  if maxSpacing != 8 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(3, nVertices, edges)
  if maxSpacing != 4 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(4, nVertices, edges)
  if maxSpacing != 1 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }
}

func TestMaxSpacingKClusters2(t *testing.T) {

  nVertices, edges := parseEdges("./clustering2.txt")
  maxSpacing := MaxSpacingKClusters(2, nVertices, edges)
  if maxSpacing != 5 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(3, nVertices, edges)
  if maxSpacing != 2 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }

  maxSpacing = MaxSpacingKClusters(4, nVertices, edges)
  if maxSpacing != 1 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", maxSpacing))
  }
}


func TestMaxClustersNSpacing(t *testing.T) {
  totalVertices, _, edges := ParseBigEdges("./clustering_big.txt", 3)
  count := MaxClustersNSpacing(3, totalVertices, edges)
  if count != 9749 {
    t.Fatal(fmt.Sprintf("Calculated max spacing k=2 %d", count))
  }
}


func TestMaxClustersNSpacing1(t *testing.T) {

  totalVertices, _, edges := ParseBigEdges("./clustering_big1.txt", 3)
  count := MaxClustersNSpacing(3, totalVertices, edges)
  if count != 2 {
    t.Fatal(fmt.Sprintf("Largest k ensuring max spacing of 3: %d. Expected %d.", count, 2))
  }
}



func TestMaxClustersNSpacing2(t *testing.T) {

  totalVertices, _, edges := ParseBigEdges("./clustering_big2.txt", 3)
  count := MaxClustersNSpacing(3, totalVertices, edges)
  if count != 2 {
    t.Fatal(fmt.Sprintf("Largest k ensuring max spacing of 3: %d. Expected %d.", count, 2))
  }

}


func TestMaxClustersNSpacin3(t *testing.T) {

  totalVertices, _, edges := ParseBigEdges("./clustering_big3.txt", 3)
  count := MaxClustersNSpacing(3, totalVertices, edges)
  if count != 6 {
    t.Fatal(fmt.Sprintf("Largest k ensuring max spacing of 3: %d. Expected %d.", count, 6))
  }

}

