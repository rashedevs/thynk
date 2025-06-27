package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
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

b, err := ioutil.ReadFile(s.Filename)
if err != nil {
    // If file does not exist, initialize empty slice and return no error
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

    return ioutil.WriteFile(s.Filename, b, 0644)
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
