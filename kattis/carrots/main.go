// Solving for Carrots
// /problems/carrots/file/statement/en/img-0001.jpg
// Photo by niznoz
// Carrots are good for you! First of all, they give you good night vision. Instead of having your lights on at home, you could eat carrots and save energy! Ethnomedically, it has also been shown that the roots of carrots can be used to treat digestive problems. In this contest, you also earn a carrot for each difficult problem you solve, or huffle-puff problems as we prefer to call them.
// You will be given the number of contestants in a hypothetical contest, the number of huffle-puff problems that people solved in the contest and a description of each contestant. Now, find the number of carrots that will be handed out during the contest.
//
// Input
// Input starts with two integers 1≤N,P≤1000 on a single line, denoting the number of contestants in the contest and the number of huffle-puff problems solved in total. Then follow N lines, each consisting of a single non-empty line in which the contestant describes him or herself. You may assume that the contestants are good at describing themselves, in a way such that an arbitrary 5-year-old with hearing problems could understand it.
//
// Output
// Output should consist of a single integer: the number of carrots that will be handed out during the contest.
//
// Sample Input 1
// 2 1
// carrots?
// bunnies
//
// Sample Output 1
// 1
//
// Sample Input 2
// 1 5
// sovl problmz
//
// Sample Output 2
// 5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	ind := strings.IndexByte(s, ' ')
	c, err := strconv.Atoi(s[ind+1:])
	if err != nil {
		log.Fatalln(err)
	}

	for scanner.Scan() {
		scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	fmt.Println(c)
}
