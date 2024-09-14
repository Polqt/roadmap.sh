package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Tracker struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Trackers []Tracker

func (trackers *Trackers) add(title string) {
	tracker := Tracker {
		Title: 			title,
		Completed: 		false,
		CompletedAt: 	nil,
		CreatedAt: 		time.Now(),
	}

	*trackers = append(*trackers, tracker)
}

func (trackers *Trackers) validateIndex(index int) error {
	if index < 0 || index >= len(*trackers) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (trackers *Trackers) delete(index int) error {
	t := *trackers

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*trackers = append(t[:index], t[index+1:]...)

	return nil
}

func (trackers *Trackers) toggle(index int) error {
		t := *trackers

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (trackers *Trackers) edit(index int, title string) error {
		t := *trackers

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

func (trackers *Trackers) print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")
	for index, t := range *trackers {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✔️"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}
