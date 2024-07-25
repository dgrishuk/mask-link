package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var sourceStr string
	fmt.Println("Пример сообщения - \"Here's my spammy page: http://hehefouls.netHAHAHA see you.\"")
	fmt.Print("Введите сообщение: ")
	sourceStr, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	fmt.Println(MaskLink(sourceStr))
}
