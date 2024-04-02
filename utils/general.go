package utils

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func GenerateNewCode(id uint32) string {
	return strconv.Itoa(rand.IntN(100)) + "UseR" + strconv.FormatUint(uint64(id), 10)
}

func PrintError(message string) {
	red := "\033[31m"
	reset := "\033[0m"
	fmt.Println(red + message + reset)
}

func PrintInfo(message string) {
	red := "\033[34m"
	reset := "\033[0m"
	fmt.Println(red + message + reset)
}
