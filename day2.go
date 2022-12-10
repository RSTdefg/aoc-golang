package main

import (
	"bufio"
	"fmt"
	"os"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	cur := 0
	mat := [3][3] int {{4, 1, 7},{8, 5, 2},{3, 9, 6}}
	for scanner.Scan(){
		line := scanner.Text()
		runeArr := []rune(line) //okay
		opp := runeArr[0] - 'A'
		me := runeArr[2] - 'X'
		op := (opp + me -1 + 3) % 3 
		fmt.Println(opp, me, mat[op][opp])
		cur += mat[op][opp]
	}
	fmt.Println(cur)
}