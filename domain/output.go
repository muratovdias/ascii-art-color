package domain

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"student/lettercolor"
)

func Print() error {
	args := os.Args[1:]
	if len(args) == 0 {
		return errors.New("EXIT STATUS : 1. ERROR: Usage: go run . [STRING] [OPTION] [LETTER || LETTERS]\nEX: go run . something --color=<color>")
	}
	if args[0] == "" {
		return nil
	}
	if CheckEmpty(args[0]) {
		count := len(args[0]) / 2
		for i := 0; i < count; i++ {
			fmt.Println()
		}
		return nil
	}
	s := CheckArgument()
	data, err := ioutil.ReadFile("standard.txt")
	if err != nil {
		return errors.New("EXIT STATUS : 2. ERROR: File is not found")
	}
	text := string(data)
	s = strings.ReplaceAll(s, "\n", "\\n")
	splitText := strings.Split(s, "\\n")
	x := MakeMap(text)
	option := "--color="
	color := ""
	if len(args) == 2 && len(args[1]) >= 8 || len(args) == 3 && len(args[1]) >= 8 {
		if option != args[1][:8] {
			return errors.New("EXIT STATUS : 3. ERROR: Usage: go run . [STRING] [OPTION] [LETTER || LETTERS]\nEX: go run . something --color=<color>")
		} else {
			a := args[1][len(option):]
			color, err = lettercolor.ChooseColor(a)
			if err != nil {
				return err
			}
		}
	} else if len(args) == 1 {
		a := "white"
		color, err = lettercolor.ChooseColor(a)
		if err != nil {
			return err
		}
	} else {
		return errors.New("EXIT STATUS : 4. ERROR: Usage: go run . [STRING] [OPTION] [LETTER || LETTERS]\nEX: go run . something --color=<color>")
	}
	var output string
	if len(args) == 3 {
		for _, w := range splitText {
			output = GetToStrSlice(w, x, color, args[2])
			fmt.Print(output)
		}
	} else if len(args) == 2 {
		for _, w := range splitText {
			output = GetToStr(w, x, color)
			fmt.Print(output)
		}
	} else if len(args) == 1 {
		color, err := lettercolor.ChooseColor("white")
		if err != nil {
			return errors.New("EXIT STATUS : 5. ERROR: Usage: go run . [STRING] [OPTION] [LETTER || LETTERS]\nEX: go run . something --color=<color>")
		}
		for _, w := range splitText {
			output = GetToStr(w, x, color)
			fmt.Print(output)
		}
	}
	return nil
}
