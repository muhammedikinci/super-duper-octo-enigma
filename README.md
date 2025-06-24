# Super Duper Octo Enigma - LSP Server

A simple Language Server Protocol (LSP) implementation in Go.

## Overview

This project implements a basic LSP server that provides language intelligence features for code editors.

## Features

- **Text Document Synchronization**
- **Diagnostics** (Basic linting/error checking)
- **Hover Information**
- **Go to Definition**
- **Code Actions**
- **Code Completion**

## Supported LSP Methods

| Method                    | Description                 |
| ------------------------- | --------------------------- |
| `initialize`              | Handshake with client       |
| `textDocument/didOpen`    | Track opening documents     |
| `textDocument/didChange`  | Track document changes      |
| `textDocument/hover`      | Provide hover information   |
| `textDocument/definition` | Go to definition support    |
| `textDocument/codeAction` | Basic code fixes/refactors  |
| `textDocument/completion` | Code completion suggestions |

## Installation

```bash
git clone https://github.com/muhammedikinci/super-duper-octo-enigma.git
cd super-duper-octo-enigma
go build
```

## Editor Configuration

Neovim

```
local client = vim.lsp.start_client {
  name = "my_lsp",
  cmd = { "{binary}" },
}

if not client then
  vim.notify "it is not working, fyi"
  return
end

vim.api.nvim_create_autocmd("FileType", {
  pattern = "markdown",
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

## Project Structure

```
├── analysis/ # State and analysis logic
├── lsp/ # LSP protocol types
├── rpc/ # Message handling
└── main.go # Server entry point
```

## Source

[Learn By Building: Language Server Protocol - TJ DeVries](https://www.youtube.com/watch?v=YsdlcQoHqPY)

[Language Server Protocol](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification)
