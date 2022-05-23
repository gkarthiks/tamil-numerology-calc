package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	alphaMap      map[string]int
	numerologySum = 0
)

func main() {
	alphaMap = map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 8, "g": 3, "h": 5, "i": 1, "j": 1, "k": 2, "l": 3, "m": 4, "n": 5, "o": 7, "p": 8, "q": 1, "r": 2, "s": 3, "t": 4, "u": 6, "v": 6, "w": 6, "x": 5, "y": 1, "z": 7}
	logrus.Println("Type the name and press enter ")
	reader := bufio.NewReader(os.Stdin)
	nameString, _ := reader.ReadString('\n')
	nameString = strings.ToLower(nameString)

	for i := 0; i < len(nameString); i++ {
		numerologySum = numerologySum + alphaMap[string(nameString[i])]
	}

	numerologyNumber := getSimplifiedNum(numerologySum)
	for numerologyNumber > 9 {
		numerologyNumber = getSimplifiedNum(numerologyNumber)
	}

	logrus.Printf("The numerology number is %d\n", numerologyNumber)
}

// getSimplifiedNum will add the adjacent numbers and returns the sum
func getSimplifiedNum(nonSimplifiedNum int) int {
	simplifiedNum := 0
	for nonSimplifiedNum > 0 {
		simplifiedNum = simplifiedNum + nonSimplifiedNum%10
		nonSimplifiedNum = nonSimplifiedNum / 10
	}
	return simplifiedNum
}
