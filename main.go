package main

import (
	"expense-tracker/storage"
	"fmt"
	"os"

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

func main() {
	// id, err := storage.Add("Payment1", -1)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Printf("Added expense with ID: %d\n", id)
	// expenses, _ := storage.List()
	// for _, e := range expenses {
	// 	fmt.Printf("[%d] %s — $%.2f\n", e.ID, e.Description, e.Amount)
	// }
}
