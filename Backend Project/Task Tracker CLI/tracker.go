package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Define status types
type Status string

const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Tracker struct {
	ID          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Trackers []Tracker

func (trackers *Trackers) getNextID() int {
	if len(*trackers) == 0 {
		return 1
	}
	return (*trackers)[len(*trackers)-1].ID + 1
}

func (trackers *Trackers) add(description string) {
	tracker := Tracker{
		ID:          trackers.getNextID(),
		Description: description,
		Status:      Todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	*trackers = append(*trackers, tracker)
}

func (trackers *Trackers) validateID(id int) (int, error) {
	for index, t := range *trackers {
		if t.ID == id {
			return index, nil
		}
	}
	err := errors.New("Invalid task ID")
	fmt.Println(err)
	return -1, err
}

func (trackers *Trackers) delete(id int) error {
	index, err := trackers.validateID(id)
	if err != nil {
		return err
	}
	*trackers = append((*trackers)[:index], (*trackers)[index+1:]...)
	return nil
}

func (trackers *Trackers) toggle(id int) error {
	index, err := trackers.validateID(id)
	if err != nil {
		return err
	}

	t := &(*trackers)[index]
	switch t.Status {
	case Todo:
		t.Status = InProgress
	case InProgress:
		t.Status = Done
	case Done:
		t.Status = Todo
	}
	t.UpdatedAt = time.Now()
	return nil
}

func (trackers *Trackers) edit(id int, description string) error {
	index, err := trackers.validateID(id)
	if err != nil {
		return err
	}
	t := &(*trackers)[index]
	t.Description = description
	t.UpdatedAt = time.Now()
	return nil
}

func (trackers *Trackers) setStatus(id int, status Status) error {
	index, err := trackers.validateID(id)
	if err != nil {
		return err
	}

	t := &(*trackers)[index]
	t.Status = status
	t.UpdatedAt = time.Now()
	return nil
}

func (trackers *Trackers) print() {
	table := table.New(os.Stdout)
	table.SetHeaders("ID", "Description", "Status", "Created At", "Updated At")

	for _, t := range *trackers {
		table.AddRow(
			strconv.Itoa(t.ID),
			t.Description,
			string(t.Status),
			t.CreatedAt.Format(time.RFC1123),
			t.UpdatedAt.Format(time.RFC1123),
		)
	}
	table.Render()
}
