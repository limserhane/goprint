package goprint

import "fmt"

func Red(tag string, description string) {
	fmt.Printf("\033[31m[%s] %s\033[0m", tag, description)
}

func Yellow(tag string, description string) {
	fmt.Printf("\033[33m[%s] %s\033[0m", tag, description)
}

func White(tag string, description string) {
	fmt.Printf("\033[37m[%s] %s\033[0m", tag, description)
}

func Green(tag string, description string) {
	fmt.Printf("\033[32m[%s] %s\033[0m", tag, description)
}

func Magenta(tag string, description string) {
	fmt.Printf("\033[35m[%s] %s\033[0m", tag, description)
}

func Blue(tag string, description string) {
	fmt.Printf("\033[34m[%s] %s\033[0m", tag, description)
}
