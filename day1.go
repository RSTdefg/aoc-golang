package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	cur := 0;
	result := make([]int, 0)
	for scanner.Scan(){
		line := scanner.Text()
		if len(line) == 0{
			result = append(result, cur)
			cur = 0
		}else{
			incr, _ := strconv.Atoi(line) //bruh theres gotta be a better way to write this
			cur += incr
		}
	}
	sort.Ints(result)
	fmt.Println(result[len(result)-1]+result[len(result)-2]+result[len(result)-3]) //what is this
}