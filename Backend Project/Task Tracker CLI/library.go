package main

import (
	"encoding/json"
	"os"
)

type Library[T any] struct {
	FileName string
}

func newLibrary[T any](fileName string) *Library[T] {
	return &Library[T]{
		FileName: fileName,
	}
}

func (l *Library[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(l.FileName, fileData, 0644)
}

func (l *Library[T]) Load(data *T) error {
	fileData, err := os.ReadFile(l.FileName)

	if err != nil {
		return err
	} 

	return json.Unmarshal(fileData, data)
}