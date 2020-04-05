package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("hello world")

	grid := [][]string{
		{"O", "O", "O", "O"},
		{"D", "O", "D", "O"},
		{"O", "O", "O", "O"},
		{"X", "D", "D", "O"},
	}

	gridMap := getSafePlaces(grid)
	paths := [][]map[string]int{}

	for i := 0; i < len(gridMap); i++ {
		findPaths(gridMap[i], &paths)
	}

	fmt.Println(gridMap)

}

func findPaths(plot map[string]int, paths *[][]map[string]int) {
	if len(*paths) == 0 {
		*paths = append(*paths, []map[string]int{plot})
		return
	}

	found := false

	for i := 0; i < len(*paths); i++ {
		for j := 0; j < len((*paths)[i]); j++ {

			temp := (*paths)[i][j]

			if plot["x"] == temp["x"] || plot["y"] == temp["y"] {

				if math.Abs(float64(plot["x"])-float64(temp["x"])) == 1 || math.Abs(float64(plot["y"])-float64(temp["y"])) == 1 {
					(*paths)[i] = append((*paths)[i], plot)
					found = true
					break
				}
			}
		}

	}
	if !found {
		*paths = append(*paths, []map[string]int{plot})
	}
}

func getSafePlaces(grid [][]string) []map[string]int {
	gridMap := []map[string]int{}

	for i, inner := range grid {
		for j, value := range inner {
			if value == "O" {
				gridMap = append(gridMap, map[string]int{"x": i, "y": j})
			}
		}
	}
	return gridMap
}
