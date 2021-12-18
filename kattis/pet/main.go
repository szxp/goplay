// Pet
// In the popular show “Dinner for Five”, five contestants compete in preparing culinary delights. Every evening one of them makes dinner and each of other four then grades it on a scale from 1 to 5. The number of points a contestant gets is equal to the sum of grades they got. The winner of the show is of course the contestant that gets the most points.
//
// Write a program that determines the winner and how many points they got.
//
// Input
// Five lines, each containing 4 integers, the grades a contestant got. The contestants are numbered 1 to 5 in the order in which their grades were given.
//
// Output
// Output on a single line the winner’s number and their points, separated by a space. The input data will guarantee that the solution is unique.
//
// Sample Input 1
// 5 4 4 5
// 5 4 4 4
// 5 5 4 4
// 5 5 5 4
// 4 4 4 5
//
// Sample Output 1
// 4 19
//
// Sample Input 2
// 4 4 3 3
// 5 4 3 5
// 5 5 2 4
// 5 5 5 1
// 4 4 4 4
//
// Sample Output 2
// 2 17

package main

import (
	"fmt"
	"log"
)

func main() {
	var g1, g2, g3, g4, win, max, sum uint8

	for i := 1; i <= 5; i++ {
		_, err := fmt.Scanln(&g1, &g2, &g3, &g4)
		if err != nil {
			log.Fatalln(err)
		}
		sum = g1 + g2 + g3 + g4
		if sum > max {
			win = uint8(i)
			max = sum
		}
	}
	fmt.Print(win, max)
}
