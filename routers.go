package main

import "fmt"

func main() {

	numRouters := int(7)
	// numLinks := int(7)
	links := [][]int{{1, 5}, {1, 3}, {2, 4}, {3, 4}, {3, 6}, {6, 7}, {4, 1}}

	fmt.Println(links)
	fmt.Println()

	impRouters := []bool{}

	for i := 1; i <= numRouters; i++ {
		// to fail each router
		temp := failRouter(links, i)

		/*
			checks if other than failed router any other router is missing from the links after failure of one router.
		*/
		if isRouterMissing(temp, numRouters, i) {
			// checks if all the links are connected
			impRouters = append(impRouters, isConnected(temp))
		} else {
			impRouters = append(impRouters, true)
		}
	}

	// fmt.Println(impRouters)

	// display output
	for i, router := range impRouters {
		if router {
			fmt.Print(i + 1)
			fmt.Print(" ")
		}
	}
}

func includes(arr []int, value int) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func isRouterMissing(links [][]int, numRouters int, curr int) bool {
	contains := false
	for i := 1; i <= numRouters; i++ {
		if i == curr {
			continue
		}

		for j := 0; j < len(links); j++ {
			if includes(links[j], i) {
				contains = true
				break
			}
			contains = false
		}

		if !contains {
			break
		}
	}
	return contains
}

func isConnected(temp [][]int) bool {
	for i := 0; i < len(temp); i++ {

		tempRouter1 := temp[i][0]
		tempRouter2 := temp[i][1]
		connFound := false

		for j := 0; j < len(temp); j++ {
			if i == j {
				continue
			}

			if includes(temp[j], tempRouter1) || includes(temp[j], tempRouter2) {
				connFound = true
				break
			}
		}

		// if not connected
		if !connFound {
			return true
		}
	}
	return false
}

func failRouter(links [][]int, router int) [][]int {
	temp := [][]int{}
	for _, link := range links {
		if !includes(link, router) {
			temp = append(temp, link)
		}
	}
	return temp
}
