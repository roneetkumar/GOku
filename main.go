package main

import "fmt"

func main() {

	alice := []int32{50, 65, 77, 90, 102}

	scores := []int32{100, 90, 90, 80, 75, 60, 60}

	aRank := []int32{}

	for _, a := range alice {
		pre := int32(0)
		rank := int32(0)
		tempRank := int32(0)

		for _, s := range scores {

			if s != pre {
				rank++
			}

			if s < a || s == a {
				tempRank = rank
				aRank = append(aRank, tempRank)
				break
			}
			pre = s
		}

		if tempRank == 0 {
			tempRank = rank + 1
			aRank = append(aRank, tempRank)
		}
	}

	fmt.Println(aRank)

}

// Alice is playing an arcade game and wants to climb to the top of the leaderboard and wants to track her ranking. The game uses Dense Ranking, so its leaderboard works like this:

// The player with the highest score is ranked number  on the leaderboard.
// Players who have equal scores receive the same ranking number, and the next player(s) receive the immediately following ranking number.
// For example, the four players on the leaderboard have high scores of , , , and . Those players will have ranks , , , and , respectively. If Alice's scores are ,  and , her rankings after each game are ,  and .

// Function Description

// Complete the climbingLeaderboard function in the editor below. It should return an integer array where each element  represents Alice's rank after the  game.

// climbingLeaderboard has the following parameter(s):

// scores: an array of integers that represent leaderboard scores
// alice: an array of integers that represent Alice's scores
