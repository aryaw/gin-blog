package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var Line = &cobra.Command{
	Use:     "bosque",
	Short:   "Bosque simple blog",
	Long:    "Bosque simple blog",
	Example: "bosque create",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			tip()
		}
	},
}

func tip() {
	printStr := `You can try using enter to get more information`
	fmt.Println(printStr)
}

func init() {
	Line.AddCommand(NewModule)
}

func Implement() {
	if err := Line.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
