package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	printer "trpo-labs/lab_2/printer"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter page count: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("there was some error due to reading string from console, err: ", err)
		return
	}

	pageCount, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		fmt.Println("failed to scan pager count! does it have integer value type?")
		return
	}

	fmt.Print("Enter list count: ")
	text, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("there was some error due to reading string from console, err: ", err)
		return
	}

	listCount, err := strconv.Atoi(strings.TrimSuffix(text, "\n"))
	if err != nil {
		fmt.Println("failed to scan list count! does it have integer value type?")
		return
	}

	fmt.Println("Result: ")
	fmt.Println(printer.Printer(pageCount, listCount))
}
