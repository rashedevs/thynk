package cmd

import (
	"flag"
	"fmt"
	"thynk/internal"
	"time"
)

func Add(args []string) {
    addFlagSet := flag.NewFlagSet("add", flag.ExitOnError)
    datePtr := addFlagSet.String("date", "", "Date for the task in YYYY-MM-DD (optional)")

    err := addFlagSet.Parse(args)
    if err != nil {
        fmt.Println("Error parsing flags:", err)
        return
    }

    if addFlagSet.NArg() < 1 {
        fmt.Println("Usage: thynk add \"task description\" [--date YYYY-MM-DD]")
        return
    }

    text := addFlagSet.Arg(0)

    date := *datePtr
    if date == "" {
        date = time.Now().Format("2006-01-02")
    } else {
        _, err := time.Parse("2006-01-02", date)
        if err != nil {
            fmt.Println("Invalid date format. Use YYYY-MM-DD")
            return
        }
    }

    store := internal.NewStorage("data.json")
    task, err := store.AddTask(text, date)
    if err != nil {
        fmt.Println("Failed to add task:", err)
        return
    }

    fmt.Printf("Added task #%d for %s: %s\n", task.ID, task.Date, task.Text)
}
