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
	nbrs := make([]int, 0)
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		var nr int
		_, err = fmt.Sscan(scan.Text(), &nr)
		if err != nil {
			log.Printf("could not scan line %v into int : %v", scan.Text(), err)
		}
		nbrs = append(nbrs, nr)
	}
	seen := map[int]bool{0: true}
	sum := 0
	for {
		for _, nr := range nbrs {
			sum += nr
			if _, ok := seen[sum]; ok {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
		}
	}
}
