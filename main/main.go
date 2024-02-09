package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
)

var Version string

func main() {
	fmt.Println("当前版本号:", Version)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
	}
}
