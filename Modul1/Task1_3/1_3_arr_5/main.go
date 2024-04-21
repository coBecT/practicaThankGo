package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var number int
	fmt.Scan(&number)
	counter := make(map[int]int)

	str := strconv.Itoa(number)
	runes := []rune(str)
	fmt.Println(len(runes))
	for i := 0; i < len(runes); i++ {
		str1 := string(runes[i])
		ch, e := strconv.Atoi(str1)
		if e != nil {
		}
		counter[ch]++
	}
	//fmt.Println(str)
	//runes := []rune(str)
	//fmt.Println(runes)
	//a, err := strconv.Atoi(string("12823"))
	//if err == nil {
	//	fmt.Println("error")
	//}
	//fmt.Println(a)
	//for i := 0; i < len(runes); i++ {
	//	counter[int(runes[i])]++
	//}
	printCounter(counter)
}
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")

}
