package main

import "strings"

func cleanInput(text string) []string{
	//string split into words and lowercased
	return strings.Fields(strings.ToLower(text))
}