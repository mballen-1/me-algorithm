package main

import (
	"bufio"
	"fmt"
	"me-algorithm/src"
	"os"
	"strings"
)

const welcomeMessage = "Enter numbers separated by a whitespace \" \" The last number will be parsed as the target number. \n " +
	"Then press Enter to proceed with the algorithm. \n eg: 1 9 5 0 20 -4 12 16 7 12 will be accepted input and target number will be 12 \n "

func main() {
	for {
		fmt.Println(welcomeMessage)
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		numsToInt, target, err := src.ParseInputString(strings.Split(text, " "))
		if err != nil {
			fmt.Printf("Error parsing input => %v \n", err)
		} else {
			printOutput(src.FindSum2(numsToInt, target))
		}
	}
}

func printOutput(answers map[int64]int64) {
	if len(answers) == 0 {
		fmt.Println("\nNo pairs found")
	} else {
		fmt.Println("\nSolution: {")
		for k, v := range answers {
			fmt.Printf(" => %v,%v \n", k, v)
		}
		fmt.Println("}")
	}
	fmt.Println()
}
