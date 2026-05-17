package main

import "fmt"

func main() {
	// //sub1
	const Author = "Abylai"
	text := "asdasd"

	fmt.Printf("Автор %s Написал:%s ", Author, text)

	fmt.Println(len(text))

	// // sub2
	message := ""

	if message == "" {
		fmt.Println("Строка пуста")
	} else {
		fmt.Println("Строка не пуста")
	}

	fmt.Println(len(message))

	// // sub3
	var input string

	fmt.Scan(&input)

	if len(input) < 5 {
		fmt.Println("Слишком короткое слово")
	} else if len(input) <= 5 || len(input) < 10 {
		fmt.Println("Нормально")
	} else {
		fmt.Println("Слишком длино")
	}

	//sub 4

	word := "Word"

	fmt.Println(string(word[0]), word[len(word)-1])

	for i, v := range word {
		fmt.Println(i, string(v))
	}

	fmt.Println(len(word))

	// sub 5

	sentence := "Arthur Arthur Arthur a a a "

	fmt.Println(len(sentence))

	count := 0

	for _, v := range sentence {
		if v == 'a' || v == 'A' {

			count++
		}
	}

	lenght := len(sentence)

	fmt.Printf("В строке %d символов и %d букв а", lenght, count)

	// sub 6

	var normal_words string
	fmt.Println()
	fmt.Scan(&normal_words)

	reversed := ""

	for _, v := range normal_words {
		reversed = string(v) + reversed

	}
	fmt.Println(reversed)

	//sub 7

	var words string

	fmt.Scan(&words)

	lenght2 := len(words)

	count_words := 1

	for _, v := range words {
		if string(v) == " " {
			count_words++
		}
	}
	count_vowel := 0
	for _, v := range words {
		if string(v) == "a" || string(v) == "e" || string(v) == "i" || string(v) == "o" || string(v) == "u" || string(v) == "y" || string(v) == "A" || string(v) == "E" || string(v) == "I" || string(v) == "O" || string(v) == "U" || string(v) == "Y" {
			count_vowel++
		}
	}

	fmt.Printf("Длина строки: %d, слов: %d, гласных: %d", lenght2, count_words, count_vowel)
}
