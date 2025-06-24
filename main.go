package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/muhammedikinci/super-duper-octo-enigma/analysis"
	"github.com/muhammedikinci/super-duper-octo-enigma/lsp"
	"github.com/muhammedikinci/super-duper-octo-enigma/rpc"
)

func main() {
	fmt.Println("hi")

	logger := getLogger("/Users/muhammedikinci/Desktop/wrk/lsp_server/log.txt")
	logger.Println("lsp server started")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	writer := os.Stdout

	state := analysis.NewState()

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, content, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("we have an error: %s\n", err)
		}
		handleMessage(logger, writer, method, content, state)
	}
}

func handleMessage(
	logger *log.Logger,
	writer io.Writer,
	method string,
	content []byte,
	state analysis.State,
) {
	logger.Printf("received with %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s\n", err)
			return
		}

		logger.Printf("connected to %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version,
		)

		msg := lsp.NewInitializeResponse(request.ID)
		writeResponse(writer, msg)

		logger.Print("Sent reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s\n", err)
			return
		}

		logger.Printf(
			"Opened: %s",
			request.Params.TextDocument.URI,
		)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.DidChangeTextDocumentNotification
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s\n", err)
			return
		}

		logger.Printf(
			"Changed: %s",
			request.Params.TextDocument.URI,
		)

		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}

	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(content, &request); err != nil {
			logger.Printf("we couldnt parse this %s\n", err)
			return
		}

		response := state.Hover(
			request.ID,
			request.Params.TextDocument.URI,
			request.Params.Position,
		)

		writeResponse(writer, response)
	}
}

func writeResponse(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}

func getLogger(filename string) *log.Logger {
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	return log.New(logFile, "[my_lsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
