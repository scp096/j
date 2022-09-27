/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/scp096/jgo/cfg"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jgo",
	Short: "A linux cmdline tool to set jump path easily and faster.",
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		shortcuts := cfg.GetShortcutsFromCfg()
		result := []string{}
		for _, shortcut := range shortcuts {
			if strings.HasPrefix(shortcut, toComplete) {
				result = append(result, shortcut)
			}
		}
		return result, cobra.ShellCompDirectiveNoFileComp
	},
}
var RootCmd = rootCmd

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
