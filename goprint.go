package goprint

import "fmt"

func Red(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[31m" + format + "\033[0m", args...)
}

func Yellow(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[33m" + format + "\033[0m", args...)
}

func White(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[37m" + format + "\033[0m", args...)
}

func Green(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[32m" + format + "\033[0m", args...)
}

func Magenta(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[35m" + format + "\033[0m", args...)
}

func Blue(format string, args ...interface{}) (string) {
	return fmt.Sprintf("\033[34m" + format + "\033[0m", args...)
}
