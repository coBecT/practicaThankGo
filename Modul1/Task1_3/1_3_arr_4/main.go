package main

import "fmt"

func main() {
	var width int
	var text string
	fmt.Scan(&text, &width)
	str := text[:width]
	str += "..."
	fmt.Println(str)

	width1 := []rune(text)
	sdfs := []byte(text)

	res := sdfs[:width]
	fmt.Println(string(res) + "...")
	fmt.Println(width1)
	//res := [0:width]string{}
	str1 := "го!"
	bytes1 := []byte(str1)
	fmt.Println(bytes1)

}
