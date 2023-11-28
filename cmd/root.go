package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain is a command-line tool to provide information about software development",
	Long: `Use Explain to remember important concepts, such as using Git commands, etc.
For example: "explain git init", "explain docker run".`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Explain! Use subcommands to get explanations.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(gitCmd)
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
