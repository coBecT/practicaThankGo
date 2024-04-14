package main

import "fmt"

func main() {
	fmt.Println("Привет! Введи язык на котором хочешь продолжить общение(например \"en\"-> English")
	s := ""
	fmt.Scan(&s)
	if s == "en" {
		fmt.Print("English")
	} else if s == "kz" {
		fmt.Print("Kazahkstan")
	} else if s == "fr" {
		fmt.Print("French")
	} else if s == "ru" || s == "rus" {
		fmt.Print("Russian")
	} else if s == "by" {
		fmt.Print("Belarussian")
	} else {
		fmt.Print("unknown таково языка в поддержке нет")
	}
}
