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
