package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

const fpath = "expenses.json"

// load: loads all expenses from the JSON storage file.
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

// save: writes the provided expenses to the JSON storage file.
func save(expenses []Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fpath, data, 0644)
}

// List: returns all saved expenses.
func List() ([]Expense, error) {
	return load()
}

// Add: creates a new expense with the given description and amount.
func Add(description string, amount float64) (int, error) {

	if description == "" {
		return 0, errors.New("description cannot be empty")
	}
	if amount <= 0 {
		return 0, errors.New("amount must be greater than zero")
	}

	expenses, err := load()
	if err != nil {
		return 0, err
	}

	id := 1
	if len(expenses) > 0 {
		id = expenses[len(expenses)-1].ID + 1
	}
	expense := Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}

	expenses = append(expenses, expense)
	return id, save(expenses)

}

// Delete: removes an expense
func Delete(id int) error {

	if id <= 0 {
		return errors.New("ID must be greater than zero")
	}

	expenses, err := load()
	if err != nil {
		return err
	}
	for i, e := range expenses {
		if e.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return save(expenses)
		}
	}
	return errors.New("expense not found")
}

// Update: changes the existing value in an expense entry
func Update(id int, description string, amount float64) error {

	if id <= 0 {
		return errors.New("ID must be greater than zero")
	}
	if amount < 0 {
		return fmt.Errorf("expense with ID %d not found", id)
	}

	expenses, err := load()
	if err != nil {
		return err
	}
	for i, e := range expenses {
		if e.ID == id {
			if description != "" {
				expenses[i].Description = description
			}
			if amount > 0 {
				expenses[i].Amount = amount
			}
			return save(expenses)
		}
	}
	return errors.New("expense not found")
}
