package main

import "fmt"

func main() {
	s := ""
	fmt.Println("Privet! Vvedi yazik \"en\"-> English")
	fmt.Scan(&s)
	switch s {
	case "en":
		fmt.Println("English")
	case "ru":
		fmt.Println("Russian")
	case "rus":
		fmt.Println("Russiam")
	case "by":
		fmt.Println("Belorussian")
	case "fr":
		fmt.Println("French")
	}

}
