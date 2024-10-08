package main

import "fmt"

type CostMatrix struct {
	values  [][]int
	rows    int
	cols    int
	cache   map[string]int
	visited [][]bool
}

func (cm *CostMatrix) IsLastRow(row int) bool {
	return row == cm.rows-1
}

func (cm *CostMatrix) IsLastCol(col int) bool {
	return col == cm.cols-1
}

func (cm *CostMatrix) IsDestination(row, col int) bool {
	return cm.IsLastRow(row) && cm.IsLastCol(col)
}

func (cm *CostMatrix) IsCached(key string) bool {
	if _, ok := cm.cache[key]; ok {
		return true
	}

	return false
}

func (cm *CostMatrix) SetCache(key string, val int) {
	cm.cache[key] = val
}

func main() {

	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	cm := &CostMatrix{
		values: m,
		rows:   len(m),
		cols:   len(m[0]),
		cache:  make(map[string]int),
		visited: [][]bool{
			{false, false, false},
			{false, false, false},
			{false, false, false},
		},
	}

	fmt.Println("matrix", m)
	var path []int
	TraversePaths(cm, 0, 0, path)
}

func TraversePaths(cm *CostMatrix, row int, col int, path []int) {

	// if cm.IsDestination() || cm.visited[row][col] {
	// 	return
	// }

	if row >= cm.rows || col >= cm.cols || cm.visited[row][col] {
		return
	}
	path = append(path, cm.values[row][col])
	if cm.IsDestination(row, col) {
		fmt.Println(path)
	}

	cm.visited[row][col] = true
	TraversePaths(cm, row+1, col, path)
	TraversePaths(cm, row, col+1, path)
	cm.visited[row][col] = false

	// key := fmt.Sprintf("(%d, %d)", row, col)
	// cm.SetCache(key, cm.values[row][col])
	// fmt.Printf("(%d, %d): %d - %d\n", row, col, cm.values[row][col], cm.cache[key])
}
