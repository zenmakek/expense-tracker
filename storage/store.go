package storage

import (
	"encoding/json"
	"errors"
	"os"
)

const fpath = "expenses.json"

func load() ([]Expense, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Expense{}, nil
		}
		return nil, err
	}
	var expenses []Expense
	err = json.Unmarshal(data, &expenses)
	return expenses, err
}

func save(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fpath, data, 0644)
}

func List() ([]Expense, error) {
	return load()
}
