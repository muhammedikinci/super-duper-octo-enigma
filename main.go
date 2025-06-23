package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/muhammedikinci/super-duper-octo-enigma/rpc"
)

func main() {
	fmt.Println("hi")

	logger := getLogger("/Users/muhammedikinci/Desktop/wrk/lsp_server/log.txt")
	logger.Println("lsp server started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[my_lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
