/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/scp096/jgo/cfg"

	"github.com/spf13/cobra"
)

// quickaddCmd represents the quickadd command
var quickaddCmd = &cobra.Command{
	Use:   "quickadd",
	Short: "Quickadd a shortcut",
	Long:  "j quickadd [shortcut]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.QuickAddRecordToCfg(args[0])
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(quickaddCmd)
}
