package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, str := range strSlice {
		strSlice[i] = strings.ToLower(str)
	}
	return strSlice
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	// scanner.Scan() 会阻塞等待输入，直到读取到下一个 token（换行符）
	// 当遇到 EOF（如按下 Ctrl+D）或错误时，它会返回 false 并退出循环
	for scanner.Scan() {
		words := cleanInput(scanner.Text())

		// 如果只输入了空格或直接按回车
		if len(words) == 0 {
			fmt.Print("Pokedex > ")
			continue
		}
		commandName := words[0]
		fmt.Printf("Your command was: %s\n", commandName)
		fmt.Print("Pokedex > ")
	}

	fmt.Println()

	// 循环结束后检查是否发生了非 EOF 的 I/O 错误
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
