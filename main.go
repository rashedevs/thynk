package main

import (
	"fmt"
	"os"
	"thynk/cmd"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: thynk <command> [arguments]")
        fmt.Println("Commands: add, today")
        return
    }

    cmdName := os.Args[1]
    args := os.Args[2:]

    switch cmdName {
    case "add":
        cmd.Add(args)
    case "today":
        cmd.Today(args)
    default:
        fmt.Println("Unknown command:", cmdName)
    }
}
