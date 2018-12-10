package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file input.txt : %v", err)
	}
	scan := bufio.NewScanner(f)
	var threes, twos int
	for scan.Scan() {
		word := make(map[rune]int)
		for _, r := range scan.Text() {
			word[r]++
		}
		var hasTwo, hasThree bool
		for _, c := range word {
			if c == 2 {
				hasTwo = true
			}
			if c == 3 {
				hasThree = true
			}
		}
		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}
	sum := threes * twos
	fmt.Println(sum)
}
