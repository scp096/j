package cfg

import (
	"bufio"
	"fmt"
	"github.com/scp096/jgo/logger"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path"
	"regexp"
	"strings"
)

var cfgFile *os.File

func getCfgFilePath() string {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	cfgPath := path.Join(u.HomeDir, ".jcfg")
	return cfgPath
}

func openCfgFile() *os.File {
	cfgPath := getCfgFilePath()
	if cfgPath == "" {
		return nil
	}

	// opend cfg file
	var err error
	cfgFile, err = os.OpenFile(cfgPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return cfgFile
}

func readCfgFile(cfgFile *os.File) ([]string, error) {
	var lines []string
	reader := bufio.NewReader(cfgFile)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				return lines, nil
			}
			return nil, err
		}
		lines = append(lines, line)
	}
}

func InitCfg() error {
	// Read cfg file
	cfgFile = openCfgFile()
	if cfgFile == nil {
		panic("open cfg file failed")
	}

	return nil
}

func UninitCfg() {
	if cfgFile != nil {
		cfgFile.Close()
	}
}

func AddRecordToCfg(param string) {
	lines, err := readCfgFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Verify input
	r, err := regexp.Compile(".+=.+")
	if err != nil {
		fmt.Println(err)
		return
	}
	isMatch := r.MatchString(param)
	if !isMatch {
		logger.ErrPrint("The input isn't valid")
		return
	}

	// Check if shortcut exists
	attrs := strings.Split(param, "=")
	shortcut := strings.TrimSpace(attrs[0])
	for _, line := range lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		if shortcut == innerShortCut {
			logger.ErrPrint(fmt.Sprintf(`Shortcut "%s" has aleady existed.`, shortcut))
			return
		}
	}

	// Save the record to file
	cfgFile.Seek(0, io.SeekEnd)
	writer := bufio.NewWriter(cfgFile)
	defer writer.Flush()
	writer.WriteString(param + "\n")
}

func QuickAddRecordToCfg(param string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	AddRecordToCfg(param + "=" + path)
}

func DeleteRecordFromCfg(param string) {
	lines, err := readCfgFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Check if shortcut exists
	nlines := len(lines)
	for index, line := range lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		if param == innerShortCut {
			lines = append(lines[:index], lines[index+1:]...)
			break
		}
	}

	if nlines == len(lines) {
		logger.ErrPrint(fmt.Sprintf(`Shortcut "%s" doesn't exist.`, param))
		return
	}

	// Save the record to file
	cfgFile.Seek(0, io.SeekStart)
	cfgFile.Truncate(0)
	writer := bufio.NewWriter(cfgFile)
	defer writer.Flush()
	for _, line := range lines {
		writer.WriteString(line + "\n")
	}
}

func GetRecordFromCfg(param string) string {
	// Read cfg file
	cfgFile := openCfgFile()
	defer cfgFile.Close()
	if cfgFile == nil {
		fmt.Println("Open cfg file failed")
		return ""
	}

	lines, err := readCfgFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	// Check if shortcut exists
	for _, line := range lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		innerPath := innerAttrs[1]
		if param == innerShortCut {
			return innerPath
		}
	}

	return ""
}

func GetShortcutsFromCfg() []string {
	var result = []string{}

	// Read cfg file
	cfgFile := openCfgFile()
	defer cfgFile.Close()
	if cfgFile == nil {
		fmt.Println("Open cfg file failed")
		return result
	}

	lines, err := readCfgFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return result
	}

	// Check if shortcut exists
	for _, line := range lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		result = append(result, innerShortCut)
	}

	return result
}

func ListRecordsFromCfg() {
	// Read cfg file
	cfgFile := openCfgFile()
	defer cfgFile.Close()
	if cfgFile == nil {
		fmt.Println("Open cfg file failed")
		return
	}

	lines, err := readCfgFile(cfgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func EditRecordFile() {
	cfgPath := getCfgFilePath()
	if cfgPath == "" {
		return
	}
	cmd := exec.Command("vim", cfgPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
