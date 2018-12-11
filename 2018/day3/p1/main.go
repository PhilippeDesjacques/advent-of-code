package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type claim struct {
	ID int
	xy
	Width int
	Height int
}

type xy struct {
	X int
	Y int
}

var fabric = map[xy]int{}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("could not open file input.txt : %v", err)
	}
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		c := claim{}
		fmt.Sscanf(scan.Text(), "#%d @ %d,%d: %dx%d", &c.ID, &c.X, &c.Y, &c.Width, &c.Height)
		for i := 0; i < c.Width; i++ {
			for j := 0; j < c.Height; j++ {
				fabric[xy{X: c.X + i, Y: c.Y + j}]++
			}
		}
	}
	cnt := 0
	for _, cn := range fabric {
		if cn > 1 {
			cnt++
		}
	}
	fmt.Println(cnt)
}
