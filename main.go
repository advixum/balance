package main

import "fmt"

func areBalanced(input string) string {
	if input == "" {
		return "Введена пустая строка"
	}
	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)
	for _, char := range input {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return "Скобки несбалансированы"
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return "Скобки сбалансированы"
	}
	return "Скобки несбалансированы"
}

func main() {
	user_input := ""
	fmt.Print("Введите тестовый случай: ")
	fmt.Scanln(&user_input)
	fmt.Println(areBalanced(user_input))
}
