package cmd

import (
	"fmt"
	"thynk/internal"
	"time"
)

func Today(args []string) {
    store := internal.NewStorage("data.json")
    tasks, err := store.GetTasks()
    if err != nil {
        fmt.Println("Failed to get tasks:", err)
        return
    }

    today := time.Now().Format("2006-01-02")
    var todaysTasks []internal.Task
    var overdueTasks []internal.Task

    for _, t := range tasks {
        if t.Date == today {
            todaysTasks = append(todaysTasks, t)
        } else if t.Date < today && !t.Completed {
            overdueTasks = append(overdueTasks, t)
        }
    }

    if len(todaysTasks) == 0 {
        fmt.Println("No tasks for today 🎉")
    } else {
        fmt.Println("📅 Tasks for today:")
        for _, t := range todaysTasks {
            fmt.Printf("- %s\n", t.Text)
        }
    }

    if len(overdueTasks) > 0 {
        fmt.Println("\n⚠️ Overdue tasks:")
        for _, t := range overdueTasks {
            fmt.Printf("- [%s] %s\n", t.Date, t.Text)
        }
    }
}
