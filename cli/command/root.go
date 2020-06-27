package command

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Short: "CLI for API",
	Example: "rli --help",
	Long: `
Welcome to Rest CLI
`,
}


func Execute(){
	if err := rootCmd.Execute();err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

