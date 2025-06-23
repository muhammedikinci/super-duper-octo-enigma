package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/muhammedikinci/super-duper-octo-enigma/lsp"
	"github.com/muhammedikinci/super-duper-octo-enigma/rpc"
)

func main() {
	fmt.Println("hi")

	logger := getLogger("/Users/muhammedikinci/Desktop/wrk/lsp_server/log.txt")
	logger.Println("lsp server started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("we have an error: %s\n", err)
		}
		handleMessage(logger, method, content)
	}
}

func handleMessage(logger *log.Logger, method string, content []byte) {
	logger.Printf("received with %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s\n", err)
		}

		logger.Printf("connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)
	}
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[my_lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
