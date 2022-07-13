package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	doc := `A linux cmdline tool to set jump path easily and faster.
Usage:
	1. To fast jump to a path:
		j [shortcut]
	2. To add a shortcut:
		j add [shortcut]=[path]
	3. To quickadd a shortcut:
		j quickadd [shortcut]
	4. To delete a shortcut:
		j delete [shortcut]
	5. To list all shortcut:
		j list
	6. To edit all shortcuts:
		j edit`
	fmt.Println(doc)
}

func errPrint(output string) {
	fmt.Println("j:", output)
	fmt.Println(`For more information, use "j --help"`)
}

func main() {
	// Parse flag args
	flag.Usage = usage
	flag.Parse()
	nargs := flag.NArg()
	if nargs < 1 {
		errPrint("At least one arg is needed.")
		return
	}
	args := flag.Args()
	cmd, params := args[0], args[1:]

	// Init cfg
	cfg := &CfgRuntime{}
	defer cfg.UninitCfg()
	err := cfg.InitCfg()
	if err != nil {
		fmt.Println(err)
		return
	}

	switch cmd {
	case "add":
		if len(params) != 1 {
			errPrint("Add method needs only one args.")
			return
		}
		cfg.AddRecordToCfg(params[0])
	case "quickadd":
		if len(params) != 1 {
			errPrint("Add method needs only one args.")
			return
		}
		cfg.QuickAddRecordToCfg(params[0])
	case "delete":
		if len(params) != 1 {
			errPrint("Delete method needs only one args.")
			return
		}
		cfg.DeleteRecordFromCfg(params[0])
	case "list":
		cfg.ListRecordsFromCfg()
	case "edit":
		cfg.EditRecord()
	default:
		shortcut := cmd
		path := cfg.GetRecordFromCfg(shortcut)
		if path == "" {
			errPrint("Unknown command or path.")
			os.Exit(-1)
		}
		fmt.Println(path)
		return
	}
}
