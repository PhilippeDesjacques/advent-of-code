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
	words := make([]string, 0)
	for scan.Scan() {
		words = append(words, scan.Text())
	}
	for i := 0; i < len(words); i++ {
		for j := 1; j < len(words); j++ {

			/*fmt.Println("i : " + strconv.Itoa(i))
			fmt.Println("j : " + strconv.Itoa(j))
			fmt.Println(words[i])
			fmt.Println(words[j])*/
			ok := compareWords(words[i], words[j])
			if ok {
				fmt.Println(findCommons(words[i], words[j]))
				return
			}
		}
	}
}

func compareWords(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	if w1 == w2 {
		return false
	}
	dif := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] == w2[i] {
			continue
		}
		dif++
		if dif > 1 {
			return false
		}
	}
	return true
}

func findCommons(w1, w2 string) string {
	c := ""
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			continue
		}
		c += string(w1[i])
	}
	return c
}