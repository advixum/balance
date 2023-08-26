package main

import (
	"testing"
)

func Test_areBalanced(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Пустая строка была обработана корректно.",
			input: "",
			want:  "Введена пустая строка",
		},
		{
			name:  "При обработке строки без скобок получен ожидаемый ответ.",
			input: "aB1",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Функция обнаружила круглые, сбалансированные скобки.",
			input: "()",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Функция обнаружила квадратные, сбалансированные скобки.",
			input: "[]",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Функция обнаружила фигурные, сбалансированные скобки.",
			input: "{}",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Функция обнаружила не закрытые, круглые скобки.",
			input: "(",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила не закрытые, квадратные скобки.",
			input: "[",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила не закрытые, фигурные скобки.",
			input: "{",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила не открытые, круглые скобки.",
			input: ")",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила не открытые, квадратные скобки.",
			input: "]",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила не открытые, фигурные скобки.",
			input: "}",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила вложенные, сбалансированные скобки.",
			input: "({}[])",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Функция обнаружила вложенные, несбалансированные скобки.",
			input: "({}[)]",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила вложенные, не закрытые скобки.",
			input: "({}[)",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Функция обнаружила вложенные, не открытые скобки.",
			input: "(}[])",
			want:  "Скобки несбалансированы",
		},
		{
			name:  "Тестовый случай 1: {()}{}()",
			input: "{()}{}()",
			want:  "Скобки сбалансированы",
		},
		{
			name:  "Тестовый случай 2: {()}(])",
			input: "{()}(])",
			want:  "Скобки несбалансированы",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := areBalanced(tt.input); got != tt.want {
				t.Errorf("areBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}