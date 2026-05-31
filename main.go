package main

import (
	"encoding/csv"
	"expense-tracker/storage"
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "expense-tracker",
	Short: "A CLI tool to manage your tools",
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		desc, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetFloat64("amount")
		category, _ := cmd.Flags().GetString("category")

		if desc == "" {
			fmt.Println("Error: description cannot be empty")
			os.Exit(1)
		}
		if amount <= 0 {
			fmt.Println("Error: amount must be greater than zero")
			os.Exit(1)
		}

		id, err := storage.Add(desc, amount, category)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Expense added successfully (ID: %d)\n", id)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")

		var expenses []storage.Expense
		var err error

		if category != "" {
			expenses, err = storage.ListByCategory(category)
		} else {
			expenses, err = storage.List()
		}

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(expenses) == 0 {
			fmt.Println("No expenses found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tDate\tDescription\tAmount\tCategory")
		fmt.Fprintln(w, "--\t----\t-----------\t------\t--------")
		for _, e := range expenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\t%s\n",
				e.ID,
				e.Date.Format("2006-01-02"),
				e.Description,
				e.Amount,
				e.Category,
			)
		}
		w.Flush()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")

		if id <= 0 {
			fmt.Println("Error: ID must be a positive integer")
			os.Exit(1)
		}

		err := storage.Delete(id)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Expense deleted successfully")
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing expense",
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		desc, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetFloat64("amount")

		if id <= 0 {
			fmt.Println("Error: ID must be a positive integer")
			os.Exit(1)
		}
		if desc == "" && amount == 0 {
			fmt.Println("Error: provide at least --description or --amount to update")
			os.Exit(1)
		}

		err := storage.Update(id, desc, amount)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("Expense updated successfully")
	},
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total expense summary",

	Run: func(cmd *cobra.Command, args []string) {
		month, _ := cmd.Flags().GetInt("month")

		if month < 0 || month > 12 {
			fmt.Println("Error: month must be between 1 and 12 (or 0 for all)")
			os.Exit(1)
		}

		expenses, err := storage.List()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		var total float64
		for _, e := range expenses {
			if month == 0 || int(e.Date.Month()) == month {
				total += e.Amount
			}
		}
		if month == 0 {
			fmt.Printf("Total Expenses: $%.2f\n", total)
		} else {
			fmt.Printf("Total Expenses for %s: $.%.2f\n", time.Month(month).String(), total)
		}

	},
}

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export expenses to a CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		fileName, _ := cmd.Flags().GetString("file")
		if fileName == "" {
			fileName = "expenses.csv"
		}

		expenses, err := storage.List()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(expenses) == 0 {
			fmt.Println("No expenses to export.")
			return
		}

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		// header row
		writer.Write([]string{"ID", "Date", "Description", "Amount", "Category"})

		// data rows
		for _, e := range expenses {
			writer.Write([]string{
				fmt.Sprintf("%d", e.ID),
				e.Date.Format("2006-01-02"),
				e.Description,
				fmt.Sprintf("%.2f", e.Amount),
				e.Category,
			})
		}

		fmt.Printf("Expenses exported to %s\n", fileName)
	},
}

func init() {

	addCmd.Flags().String("description", "", "Description of the expense")
	addCmd.Flags().Float64("amount", 0, "Amount of the expense")
	addCmd.Flags().String("category", "", "Category of the expense (e.g. food, travel)")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")

	deleteCmd.Flags().Int("id", 0, "ID of the expense to delete")
	deleteCmd.MarkFlagRequired("id")

	listCmd.Flags().String("category", "", "Filter by category")

	updateCmd.Flags().Int("id", 0, "ID of the expense to update")
	updateCmd.Flags().String("description", "", "New description")
	updateCmd.Flags().Float64("amount", 0, "New amount")
	updateCmd.MarkFlagRequired("id")

	summaryCmd.Flags().Int("month", 0, "Month number (1-12), 0 = all")

	exportCmd.Flags().String("file", "expenses.csv", "Output file name")

	rootCmd.AddCommand(exportCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(summaryCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
