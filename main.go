package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	alphaMap map[string]int
	numerologySum = 0
)

func init() {
	alphaMap = make(map[string]int)

	alphaMap["a"]=1
	alphaMap["b"]=2
	alphaMap["c"]=3
	alphaMap["d"]=4
	alphaMap["e"]=5
	alphaMap["f"]=8
	alphaMap["g"]=3
	alphaMap["h"]=5
	alphaMap["i"]=1
	alphaMap["j"]=1
	alphaMap["k"]=2
	alphaMap["l"]=3
	alphaMap["m"]=4
	alphaMap["n"]=5
	alphaMap["o"]=7
	alphaMap["p"]=8
	alphaMap["q"]=1
	alphaMap["r"]=2
	alphaMap["s"]=3
	alphaMap["t"]=4
	alphaMap["u"]=6
	alphaMap["v"]=6
	alphaMap["w"]=6
	alphaMap["x"]=5
	alphaMap["y"]=1
	alphaMap["z"]=7
}

func main() {
	logrus.Println("Type the name and press enter ")
	reader := bufio.NewReader(os.Stdin)
	nameString, _ := reader.ReadString('\n')
	nameString = strings.ToLower(nameString)

	for i:=0; i< len(nameString); i++ {
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
		simplifiedNum = simplifiedNum + nonSimplifiedNum % 10
		nonSimplifiedNum = nonSimplifiedNum/10
	}
	return simplifiedNum
}