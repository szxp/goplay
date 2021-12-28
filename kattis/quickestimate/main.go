// Quick Estimates
// /problems/quickestimate/file/statement/en/img-0001.jpg
// Photo by Simon A. Eugster
// Let’s face it... you are not that handy. When you need to make a major home repair, you often need to hire someone to help. When they come for the first visit, they make an estimate of the cost. Here they must be careful: if they overestimate the cost, it might scare you off, but if they underestimate, the work might not be worth their time.
//
// Because the worker is so careful, it can take a long time for them to produce the estimate. But that’s frustrating — when you ask for an estimate, you really are asking for the magnitude of the cost. Will this be $10 or $100 or $1000? That’s all you really want to know on a first visit.
//
// Please help the worker make the type of estimate you desire. Write a program that, given the worker’s estimate, reports just the magnitude of the cost — the number of digits needed to represent the estimate.
//
// Input
// Input begins with a line containing an integer N (1≤N≤100). The next N lines each contain one estimated cost, which is an integer between 0 and 10100. (Some of the workers overcharge quite a bit.)
//
// Output
// For each estimated cost, output the number of digits required to represent it.
//
// Sample Input 1
// 5
// 314
// 1
// 5926
// 5
// 35897
//
// Sample Output 1
// 3
// 1
// 4
// 1
// 5
//
// Sample Input 2
// 3
// 0
// 10
// 100
//
// Sample Output 2
// 1
// 2
// 3

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	err := solve()
	if err != nil {
		log.Fatalln(err)
	}
}

func solve() error {
	w := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(os.Stdin)

	// skip the first line
	scanner.Scan()

	var b []byte
	for scanner.Scan() {
		b = scanner.Bytes()
		fmt.Fprintln(w, len(b))
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return w.Flush()
}


