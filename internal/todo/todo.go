package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type todoItem struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []todoItem

func (t *Todos) Add(task string) {
	todo := todoItem{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(id int) error {
	tasksList := *t
	if id <= 0 || id > len(tasksList) {
		return errors.New("invaild id")
	}

	tasksList[id-1].CompletedAt = time.Now()
	tasksList[id-1].Done = true

	return nil
}

func (t *Todos) UnComplete(id int) error {
	tasksList := *t
	if id <= 0 || id > len(tasksList) {
		return errors.New("invaild id")
	}

	tasksList[id-1].CompletedAt = time.Time{}
	tasksList[id-1].Done = false

	return nil
}

func (t *Todos) Delete(id int) error {
	tasksList := *t
	if id <= 0 || id > len(tasksList) {
		return errors.New("invaild id")
	}

	*t = append(tasksList[:id-1], tasksList...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	for index, item := range *t {
		fmt.Printf("%d - %s\n", index+1, item.Task)
	}
}
