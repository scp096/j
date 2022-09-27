/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"github.com/scp096/jgo/cfg"
	"github.com/scp096/jgo/cmd"
	"github.com/scp096/jgo/logger"
	"os"
)

func find(source []string, value string) bool {
	for _, item := range source {
		if item == value {
			return true
		}
	}
	return false
}

func main() {
	// Init cfg
	err := cfg.InitCfg()
	if err != nil {
		panic(err)
	}
	defer cfg.UninitCfg()
	os.Args[0] = "j"
	if !find(os.Args[1:], "__completeNoDesc") &&
		!find(os.Args[1:], "__complete") { // not complete script call
		subCmd, _, err := cmd.RootCmd.Find(os.Args[1:])
		// not found
		if err != nil || subCmd == nil {
			shortcut := os.Args[1]
			path := cfg.GetRecordFromCfg(shortcut)
			if path == "" {
				logger.ErrPrint("Unknown command or path.")
				os.Exit(-1)
			}
			fmt.Println(path) // print to bash j function to cd
			return
		}
	}

	cmd.Execute()
}
