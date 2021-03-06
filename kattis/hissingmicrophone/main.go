// Hissing Microphone
// A known problem with some microphones is the “hissing s”. That is, sometimes the sound of the letter s is particularly pronounced; it stands out from the rest of the word in an unpleasant way.
//
// Of particular annoyance are words that contain the letter s twice in a row. Words like amiss, kiss, mississippi and even hiss itself.
//
// Input
// The input contains a single string on a single line. This string consists of only lowercase letters (no spaces) and has between 1 and 30 characters.
//
// Output
// Output a single line. If the input string contains two consecutive occurrences of the letter s, then output hiss. Otherwise, output no hiss.
//
// Sample Input 1
// amiss
//
// Sample Output 1
// hiss
//
// Sample Input 2
// octopuses
//
// Sample Output 2
// no hiss
//
// Sample Input 3
// hiss
//
// Sample Output 3
// hiss

package main

import (
	"io"
	"log"
	"os"
)

func main() {
	b := make([]byte, 32)
	var found bool
	for {
		n, err := os.Stdin.Read(b)

		for i, c := range b[:n] {
			if c == 's' && i < n-1 && b[i+1] == 's' {
				found = true
				break
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}
	}

	if found {
		os.Stdout.Write([]byte("hiss"))
	} else {
		os.Stdout.Write([]byte("no hiss"))
	}
}
