package domain

import "strings"

func GetToStr(s string, x map[rune]string, color string) string {
	temp := [8]string{}
	if s == "" {
		return "\n"
	}
	for _, char := range s {
		for index, value := range strings.Split(x[char], "\n") {
			temp[index] += color + value
		}
	}
	result := ""
	for _, value := range temp {
		result += value + "\n"
	}
	return result
}
