# Worst-GO-HttpServer

A simple HTTP server implementation in Go.

## Features

- Single file implementation
- Logs all incoming requests
- Responds with the requested path

## Running the server

```bash
# Default port (8080)
go run main.go

# Custom port
PORT=3000 go run main.go
```

## Endpoints

- `GET /` - Returns "Hello, you've requested: /"
- Any path will return "Hello, you've requested: [path]" 