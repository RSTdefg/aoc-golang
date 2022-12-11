package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	cur := 0
	index := 0
	s := map[rune]int{}
	for scanner.Scan(){
		line := scanner.Text()
		l := map[rune]bool{}
		fmt.Println(line)
		for _, c := range line{
			if _, ok := l[c]; !ok{
				s[c] +=1
				l[c] = true
			}
		}
		fmt.Println(index)
		if (index+1) % 3 == 0{
			for c, count := range s{
				if count == 3{
					if c >= 97{
						cur += int(c- 'a') + 1
					}else{
						cur += int(c - 'A') + 27
					}
				}
			}	
			s = map[rune]int{}
		}
		index += 1
	}
	fmt.Println(cur)
}