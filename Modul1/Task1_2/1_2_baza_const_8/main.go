package main

import (
	"fmt"
	"time"
)

func main() {
	s := ""
	times := 0
	fmt.Scan(&s, &times)
	var result string
	for i := 1; i <= times; i++ {
		fmt.Printf("Строка %s, выведена: %d раз\n", s, i)
		result += s
	}

	fmt.Println(result)
	switch times {
	case 1:
		fmt.Print(result)
	case 2:
		fmt.Print(result)
	case 3:
		fmt.Printf("Ден: %v\n", time.Now().Day())
		fallthrough
	case 4:
		fmt.Printf("Час - %d", time.Now().Hour())
	}
}
