package main

import (
	"antiplagiarism/core"
	"antiplagiarism/server"
	"fmt"
)

func main() {
	core := core.NewCore()
	handlers := server.NewHandlers(core)
	server := server.NewHTTPServer(handlers)

	if err := server.Run(); err != nil {
		fmt.Println("Server error:", err)
		return
	}
}