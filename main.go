package main

import (
	"fmt"
	"os"
	"strconv"
	"thynk/cmd"
	"thynk/internal"
	// "storage"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: thynk <command> [arguments]")
        fmt.Println("Commands: add, today, list, complete")
        return
    }

    storage := internal.NewStorage("data.json")
    cmdName := os.Args[1]
    args := os.Args[2:]

    switch cmdName {
    case "add":
        cmd.Add(storage, args)
    case "today":
        cmd.Today(storage, args)
    case "list":
        storage.ListTasks()   
    case "complete":
        if len(os.Args) < 3 {
            fmt.Println("Usage: thynk complete <task-id>")
            return
        }
        id, err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println("Invalid task ID")
            return
        }
        err = storage.CompleteTask(id)
        if err != nil {
            fmt.Println("Error:", err)
        } else {
            fmt.Printf("✅ Task %d marked as completed.\n", id)
        }

    default:
        fmt.Println("Unknown command:", cmdName)
    }
}
