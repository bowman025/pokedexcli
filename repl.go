package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowText := strings.ToLower(text)
	return strings.Fields(lowText)
}
