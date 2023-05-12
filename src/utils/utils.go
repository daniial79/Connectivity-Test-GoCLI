package utils

import "strings"

func IsExit(input string) bool {
	lowerInput := strings.ToLower(input)
	return strings.Compare(lowerInput, "exit") == 0
}

func CommandIsValid(command string) bool {
	lowerCommand := strings.ToLower(command)
	return strings.Compare(lowerCommand, "checkinfyconnection") == 0
}
