package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func Execute() {
	var echoTimes int

	rootCmd := &cobra.Command{
		Use:     "app",
		Short:   "A brief description of your application",
		Version: "0.0.1",
		Long:    `A longer description that spans multiple lines and likely contains`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Cobra!")
		},
	}

	cmdPrint := &cobra.Command{
		Use:   "print",
		Short: "Prints a message",
		Long:  `Prints a message`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Print: %s\n", strings.Join(args, " "))
		},
	}

	cmdEcho := &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	cmdTimes := &cobra.Command{
		Use:   "times [number of times to repeat]",
		Short: "Prints a message",
		Long:  `Prints a message`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
		},
	}

	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")

	rootCmd.AddCommand(cmdPrint, cmdEcho)
	cmdEcho.AddCommand(cmdTimes)

	rootCmd.Execute()
}
