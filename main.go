package main

import (
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
		id, err := storage.Add(desc, amount)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Expense added! (ID: %d\n)", id)

	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Run: func(cmd *cobra.Command, args []string) {
		expenses, err := storage.List()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(expenses) == 0 {
			fmt.Println("No expenses found.")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tDate\tDescription\tAmount")
		fmt.Fprintln(w, "--\t----\t-----------\t------")
		for _, e := range expenses {
			fmt.Fprintf(w, "%d\t%s\t%s\t$%.2f\n",
				e.ID,
				e.Date.Format("2006-01-02"),
				e.Description,
				e.Amount,
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
			fmt.Printf("Total Expenses: Rs%.2f\n", total)
		} else {
			fmt.Printf("Total Expenses for %s: Rs.%.2f\n", time.Month(month).String(), total)
		}

	},
}

func init() {

	addCmd.Flags().String("description", "", "Description of the expense")
	addCmd.Flags().Float64("amount", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")

	deleteCmd.Flags().Int("id", 0, "ID of the expense to delete")
	deleteCmd.MarkFlagRequired("id")

	updateCmd.Flags().Int("id", 0, "ID of the expense to update")
	updateCmd.Flags().String("description", "", "New description")
	updateCmd.Flags().Float64("amount", 0, "New amount")
	updateCmd.MarkFlagRequired("id")

	summaryCmd.Flags().Int("month", 0, "Month number (1-12), 0 = all")

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
