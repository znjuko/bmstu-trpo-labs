package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func PlayHanoi(n int, from, to, total string) {
	if n <= 0 {
		return
	}

	PlayHanoi(n-1, from, total, to)
	fmt.Println(fmt.Sprintf("moving disk '%v' from %s to %s", n, from, to))
	PlayHanoi(n-1, total, to, from)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter disk count: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("there was some error due to reading string from console, err: ", err)
		return
	}

	diskCount, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		fmt.Println("failed to scan disk count! does it have integer value type?")
		return
	}

	fmt.Println("Starting playing the Hanoi:")
	PlayHanoi(diskCount, "a", "b", "c")
}
