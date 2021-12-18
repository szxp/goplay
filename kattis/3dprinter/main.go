// 3D Printed Statues
// /problems/3dprinter/file/statement/en/img-0001.jpg
// Picture by Ariosvaldo Gonzáfoles, cc-by
// You have a single 3D printer, and would like to use it to produce n statues. However, printing the statues one by one on the 3D printer takes a long time, so it may be more time-efficient to first use the 3D printer to print a new printer. That new printer may then in turn be used to print statues or even more printers. Print jobs take a full day, and every day you can choose for each printer in your possession to have it print a statue, or to have it 3D print a new printer (which becomes available for use the next day).
//
// What is the minimum possible number of days needed to print at least n statues?
//
// Input
// The input contains a single integer n (1≤n≤10000), the number of statues you need to print.
//
// Output
// Output a single integer, the minimum number of days needed to print at least n statues.
//
// Sample Input 1
// 1
//
// Sample Output 1
// 1
//
// Sample Input 2
// 5
//
// Sample Output 2
// 4

package main

import (
	"fmt"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatalln(err)
	}

	i := 0
	for ; i < 15; i++ {
		//fmt.Println(n, 1<<i)
		if n <= 1<<i {
			break
		}
	}

	fmt.Print(i + 1)
}
