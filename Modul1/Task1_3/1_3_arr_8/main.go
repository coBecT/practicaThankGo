package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Today I learned
// Высшее учебное заведение
// Кот обладает талантом
// Ар 2 Ди #2
func main() {
	phrase := readString()

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...
	str := strings.Fields(phrase)
	//fmt.Println(str)
	abbr := ""
	//outStr := make([]string, len(str))
	//out := make([]string, len(str))
	for _, value := range str {
		ru := []rune(value)
		if unicode.IsLetter(ru[0]) {
			abbr += string(unicode.ToUpper(ru[0]))
		}
		//outStr = append(outStr, unicode.ToUpper(value))
		//fmt.Println(index, value)
	}
	fmt.Println(abbr)
	//fmt.Println(out)
	//fmt.Printf("Fields are: %q", str)
	//str2 := make([]rune, len(str))
	//for i, val := range str {
	//	str[i] = unicode.ToUpper(rune())
	//}
	//
	//fmt.Println(string(abbr))
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}

//package main
//
//import (
//"bufio"
//"fmt"
//"os"
//"strings"
//"unicode"
//)
//
//func main() {
//	phrase := readString()
//
//	// 1. Разбейте фразу на слова, используя `strings.Fields()`
//	// 2. Возьмите первую букву каждого слова и приведите
//	//    ее к верхнему регистру через `unicode.ToUpper()`
//	// 3. Если слово начинается не с буквы, игнорируйте его
//	//    проверяйте через `unicode.IsLetter()`
//	// 4. Составьте слово из получившихся букв и запишите его
//	//    в переменную `abbr`
//	// ...
//
//	fmt.Println(string(abbr))
//}
//
//// ┌─────────────────────────────────┐
//// │ не меняйте код ниже этой строки │
//// └─────────────────────────────────┘
//
//// readString читает строку из `os.Stdin` и возвращает ее
//func readString() string {
//	rdr := bufio.NewReader(os.Stdin)
//	str, _ := rdr.ReadString('\n')
//	return str
//}
