package domain

import (
	"strings"
	"student/lettercolor"
)

func GetToStrSlice(s string, x map[rune]string, color, letter string) string {
	temp := [8]string{}
	if s == "" {
		return "\n"
	}
	standardColor, err := lettercolor.ChooseColor("white")
	if err != nil {
		return "ERROR"
	}
	flag := false
	var l rune
	for _, char := range s {
		for _, letter := range letter {
			if char == letter {
				flag = true
				l = letter
			}
		}
		if flag {
			for index, value := range strings.Split(x[l], "\n") {
				temp[index] += color + value + standardColor
			}
		} else {
			for index, value := range strings.Split(x[char], "\n") {
				temp[index] += value
			}
		}
		flag = false
	}
	result := ""
	for _, value := range temp {
		result += value + "\n"
	}
	return result
}
