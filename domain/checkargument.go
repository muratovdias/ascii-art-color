package domain

import (
	"os"
)

func CheckArgument() string {
	args := os.Args[1:]
	if len(args) != 1 && len(args) != 2 && len(args) != 3 {
		return "Exit status: 6. Usage: go run . [STRING] [OPTION] [LETTER || LETTERS]\nEX: go run . something --color=<color>"
	}
	for _, word := range args[0] {
		if word != 10 && (word < 32 || word > 126) {
			return "This is not ASCII symbol"
		}
	}
	return args[0]
}
