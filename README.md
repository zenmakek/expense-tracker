# Expense Tracker CLI (WIP)

A fast, lightweight command-line expense manager written in Go. Sole purpose of learning GoLang.

---

## Features

- Add expenses with a description and amount
- Update or delete existing expenses
- List all recorded expenses in a clean table
- View total expense summary
- Filter summary by a specific month (WIP)
- Assign categories to expenses and filter by them (WIP)
- Set monthly budgets with overspend warnings (WIP)
- Export expenses to a CSV file (WIP)

---

## Tech Stack

| Layer | Technology |
|---|---|
| Language | Go 1.22+ |
| CLI Framework | [Cobra](https://github.com/spf13/cobra) |
| Storage | JSON (flat file) |
| Export | CSV via `encoding/csv` |

---

## Project Structure

```
expense-tracker/
├── go.mod
├── go.sum
├── main.go              # Entry point
├── expenses.json        # Auto-generated data file
└── storage/
    ├── expense.go       # Expense struct definition
    └── store.go         # Load, save, add, delete, update logic
```

---

## Getting Started

### Prerequisites

- [Go 1.22+](https://go.dev/dl/) installed on your machine

### Installation

```bash
# Clone the repository
git clone https://github.com/your-username/expense-tracker.git
cd expense-tracker

# Install dependencies
go mod tidy

# Build the binary
go build -o expense-tracker .
```

### (Optional) Add to PATH

```bash
# Linux / macOS
mv expense-tracker /usr/local/bin/

# Now you can run it from anywhere
expense-tracker --help
```

---

## Usage

### Add an expense

```bash
$ expense-tracker add --description "Lunch" --amount 20
Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
Expense added successfully (ID: 2)
```

### List all expenses

```bash
$ expense-tracker list
ID   Date         Description   Amount
1    2024-08-06   Lunch         $20.00
2    2024-08-06   Dinner        $10.00
```

### Update an expense

```bash
$ expense-tracker update --id 1 --description "Lunch with team" --amount 35
Expense updated successfully
```

### Delete an expense

```bash
$ expense-tracker delete --id 2
Expense deleted successfully
```

### View summary

```bash
$ expense-tracker summary
Total expenses: $35.00
```

### View summary for a specific month

```bash
$ expense-tracker summary --month 8
Total expenses for August: $35.00
```

---

## Bonus Features

### Filter by category

```bash
$ expense-tracker add --description "Netflix" --amount 15 --category entertainment
$ expense-tracker list --category entertainment
```

### Set a monthly budget

```bash
$ expense-tracker budget --month 8 --amount 500
Budget set: $500.00 for August

# Automatically warns you when you exceed it:
⚠ Warning: You have exceeded your August budget of $500.00 (spent $520.00)
```

### Export to CSV

```bash
$ expense-tracker export --file expenses.csv
Expenses exported to expenses.csv
```

---

## Error Handling

The CLI handles common edge cases gracefully:

- Negative or zero amounts are rejected
- Deleting or updating a non-existent ID shows a clear error
- Invalid month values (outside 1–12) are caught
- Missing required flags prompt a usage hint

---

## Data Storage

All expenses are stored locally in `expenses.json` in the project root. The file is created automatically on first use. You can back it up, move it, or inspect it directly:

```json
[
  {
    "id": 1,
    "description": "Lunch",
    "amount": 20,
    "date": "2024-08-06T13:45:00Z"
  }
]
```

---

## Learning Context

This project was built as a guided learning exercise to get hands-on with Go fundamentals — including structs, JSON serialization, file I/O, slices, error handling, and CLI design with Cobra. Each module of the codebase maps to a specific Go concept.

---

## Contributing

Pull requests are welcome. For major changes, open an issue first to discuss what you'd like to change.

---

## License

[MIT](LICENSE)