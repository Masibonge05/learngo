package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your full name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter current year: ")
	year, _ := reader.ReadString('\n')

	cleanName := strings.Replace(name, "\n", "", -1)
	cleanYear := strings.Replace(year, "\n", "", -1)

	

	fmt.Printf("clean name: %v \n", cleanName)
	fmt.Printf("clean year: %v \n", cleanYear)

}
