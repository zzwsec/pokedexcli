package main

import "strings"

func cleanInput(text string) []string {
	strSlice := strings.Fields(text)
	for i, str := range strSlice {
		strSlice[i] = strings.ToLower(str)
	}
	return strSlice
}
