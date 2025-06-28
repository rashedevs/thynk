package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	// "io/ioutil"
	"os"
	"sync"
)

type Task struct {
    ID        int    `json:"id"`
    Text      string `json:"text"`
    Date      string `json:"date"`      
    Completed bool   `json:"completed"`
}

type Storage struct {
    Filename string
    mu       sync.Mutex
    Tasks    []Task
}

func NewStorage(filename string) *Storage {
    return &Storage{Filename: filename}
}

func (s *Storage) load() error {
    s.mu.Lock()
    defer s.mu.Unlock()

    if _, err := os.Stat(s.Filename); errors.Is(err, os.ErrNotExist) {
        s.Tasks = []Task{}
        return nil
    }

b, err := os.ReadFile(s.Filename)
if err != nil {
    if errors.Is(err, os.ErrNotExist) {
        s.Tasks = []Task{}
        return nil
    }
    return err
}

if len(b) == 0 {
    s.Tasks = []Task{}
    return nil
}

return json.Unmarshal(b, &s.Tasks)

}

func (s *Storage) save() error {
    s.mu.Lock()
    defer s.mu.Unlock()

    b, err := json.MarshalIndent(s.Tasks, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(s.Filename, b, 0644)
}

func (s *Storage) AddTask(text, date string) (Task, error) {
    if err := s.load(); err != nil {
        return Task{}, err
    }

    newID := 1
    for _, t := range s.Tasks {
        if t.ID >= newID {
            newID = t.ID + 1
        }
    }

    task := Task{
        ID:   newID,
        Text: text,
        Date: date,
    }
    s.Tasks = append(s.Tasks, task)

    if err := s.save(); err != nil {
        return Task{}, err
    }

    return task, nil
}

func (s *Storage) GetTasks() ([]Task, error) {
    if err := s.load(); err != nil {
        return nil, err
    }
    return s.Tasks, nil
}

func (s *Storage) CompleteTask(id int) error {
    if err := s.load(); err != nil {
            return err
        }

    for i, task := range s.Tasks {
        if task.ID == id {
            if s.Tasks[i].Completed {
                return fmt.Errorf("task %d is already marked as complete", id)
            }
            s.Tasks[i].Completed = true
            return s.save()
        }
    }
    return fmt.Errorf("task with ID %d not found", id)
}

// func (s *Storage) ListTasks() {
//     if err := s.load(); err != nil {
//         fmt.Println("Failed to load tasks:", err)
//         return
//     }
//     if len(s.Tasks) == 0 {
//         fmt.Println("No tasks found.")
//         return
//     }

//     for _, task := range s.Tasks {
//         status := "❌"
//         if task.Completed {
//             status = "✅"
//         }

//         fmt.Printf("%s %d. %s", status, task.ID, task.Text)
//         if task.Date != "" {
//             fmt.Printf(" (Due: %s)", task.Date)
//         }
//         fmt.Println()
//     }
// }

func (s *Storage) ListTasks() error {
    if err := s.load(); err != nil {
        return err
    }

    if len(s.Tasks) == 0 {
        fmt.Println("No tasks found.")
        return nil
    }

    for _, task := range s.Tasks {
        status := "❌"
        if task.Completed {
            status = "✅"
        }

        fmt.Printf("%s %d. %s", status, task.ID, task.Text)
        if task.Date != "" && !task.Completed {
            fmt.Printf(" (Due: %s)", task.Date)
        }
        fmt.Println()
    }

    return nil
}
