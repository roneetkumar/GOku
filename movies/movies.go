package main

import "fmt"

func main() {

	// movie_duration: [90, 85, 75, 60, 120, 150, 125]
	// d: 250

	moviesDuration := []int{90, 85, 75, 60, 120, 150, 125}

	flightDuration := 220

	durationMap := []map[string]int{}

	for i, n := range moviesDuration {

		for j, m := range moviesDuration {

			if i == j {
				continue
			}

			durationMap = append(durationMap, map[string]int{"movie1": n, "movie2": m})

		}
	}

	// selectedMovies := []map[string]int{}

	max := 0
	temp := map[string]int{}

	for _, dur := range durationMap {

		if dur["movie1"]+dur["movie2"] > flightDuration {
			continue
		}
		if dur["movie1"]+dur["movie2"] > max {
			temp = dur
			// selectedMovies = append(selectedMovies, temp)
		}

	}

	fmt.Println(temp)

}
