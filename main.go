package main

import (
	"expense-tracker/storage"
	"fmt"
	"os"
	"text/tabwriter"

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
