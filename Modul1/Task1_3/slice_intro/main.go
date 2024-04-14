package main

import "fmt"

func main() {
	//slice из 3 строк
	s := make([]string, 3)
	fmt.Println(s)
	//Шаблон %#v возвращает «внутреннее представление» значения, примерно как repr() в питоне
	fmt.Printf("empty %#v\n", s)
	s = append(s, "f")
	fmt.Printf("empty %#v\n", s)
	s = append(s, "d", "e")
	fmt.Printf("empty %#v\n", s)
	slc := []string{"bobro1", "dobree2"}
	fmt.Println(slc)
}
