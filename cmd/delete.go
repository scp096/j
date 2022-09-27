/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/scp096/jgo/cfg"
	"github.com/spf13/cobra"
	"strings"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a shortcut",
	Long:  "j delete [shortcut]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg.DeleteRecordFromCfg(args[0])
	},
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

func init() {
	rootCmd.AddCommand(deleteCmd)
}
