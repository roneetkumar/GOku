package main

import (
	"fmt"
	"math"
)

func main() {

	grid := [][]int{
		{1, 1, 1, 1},
		{1, 0, 1, 0},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}

	gridMap := getLocationMap(grid)

	clusters := [][]map[string]int{}

	for i := 0; i < len(gridMap); i++ {
		checkInCluster(gridMap[i], &clusters)
	}

	fmt.Println("Total number of clusters are : ", len(clusters))

	for i, cluster := range clusters {
		fmt.Println("Cluster ", i+1, ": ", cluster)
	}
}

/*
checkInCluster check if plot is part of the one of the clusters; if yes the plot
gets added to the cluster else it get added as new cluster
*/
func checkInCluster(plot map[string]int, clusters *[][]map[string]int) {

	if len(*clusters) == 0 {
		*clusters = append(*clusters, []map[string]int{plot})
	}

	for i := 0; i < len(*clusters); i++ {
		var found bool
		for j := 0; j < len((*clusters)[i]); j++ {
			found = false

			temp := (*clusters)[i][j]

			if plot["x"] == temp["x"] || plot["y"] == temp["y"] {

				if math.Abs(float64(plot["x"])-float64(temp["x"])) == 1 || math.Abs(float64(plot["y"])-float64(temp["y"])) == 1 {

					(*clusters)[i] = append((*clusters)[i], plot)

				}
			} else {
				found = true
				continue
			}
		}

		if found {
			*clusters = append(*clusters, []map[string]int{plot})
			break
		}
	}
}

// get all the location map
func getLocationMap(grid [][]int) []map[string]int {

	gridMap := []map[string]int{}

	for i, inner := range grid {
		for j, value := range inner {
			if value == 1 {
				gridMap = append(gridMap, map[string]int{"x": i, "y": j})
			}
		}
	}
	return gridMap
}
