package main

import (
	"fmt"
)

func main() {

	//current state of houses
	houses := [4]int{1, 0, 0, 1}

	fmt.Println("Current State : ", houses)

	// final state of houses
	finalState := changeState(houses, 3)

	fmt.Println("Final State : ", finalState)
}

func changeState(houses [4]int, days int) [4]int {
	// temporary house to change the state next day
	tempHouse := houses

	// loop the number of days
	for i := 0; i < days; i++ {
		// current position
		cur := 0

		// dynamic last position
		last := len(houses) - 1

		// loop to check each house's state
		for cur <= (last) {
			// condiation to check if the house is either first or last
			if cur == 0 || cur == last {
				// changing curr house based on adjecent houses
				if houses[1] == 0 || houses[last-1] == 0 {
					tempHouse[cur] = 0
				} else {
					tempHouse[cur] = 1
				}
			} else {
				// changing curr house based on adjecent houses if curr house is not first or last
				if houses[cur-1] == houses[cur+1] {
					tempHouse[cur] = 0
				} else {
					tempHouse[cur] = 1
				}
			}
			// increment curr for next house
			cur++
		}

		// change the state of the houses next day
		houses = tempHouse
		fmt.Println(houses)
	}

	// return the final state after (n) no. of days
	return houses
}
