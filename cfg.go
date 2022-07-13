package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"path"
	"regexp"
	"strings"
)

type CfgRuntime struct {
	cfgFile *os.File
	lines   []string
}

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
	cfgFile, err := os.OpenFile(cfgPath, os.O_RDWR|os.O_CREATE, os.ModePerm)
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

func (c *CfgRuntime) InitCfg() error {
	// Read cfg file
	c.cfgFile = openCfgFile()
	if c.cfgFile == nil {
		fmt.Println("Open cfg file failed")
		return errors.New("init cfg failed")
	}

	var err error
	c.lines, err = readCfgFile(c.cfgFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (c *CfgRuntime) UninitCfg() {
	c.cfgFile.Close()
}

func (c *CfgRuntime) AddRecordToCfg(param string) {
	// Verify input
	r, err := regexp.Compile(".+=.+")
	if err != nil {
		fmt.Println(err)
		return
	}
	isMatch := r.MatchString(param)
	if !isMatch {
		errPrint("The input isn't valid")
		return
	}

	// Check if shortcut exists
	attrs := strings.Split(param, "=")
	shortcut := strings.TrimSpace(attrs[0])
	for _, line := range c.lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		if shortcut == innerShortCut {
			errPrint(fmt.Sprintf(`Shortcut "%s" has aleady existed.`, shortcut))
			return
		}
	}

	// Save the record to file
	c.cfgFile.Seek(0, io.SeekEnd)
	writer := bufio.NewWriter(c.cfgFile)
	defer writer.Flush()
	writer.WriteString(param + "\n")
}

func (c *CfgRuntime) QuickAddRecordToCfg(param string) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	c.AddRecordToCfg(param + "=" + path)
}

func (c *CfgRuntime) DeleteRecordFromCfg(param string) {
	// Check if shortcut exists
	nlines := len(c.lines)
	for index, line := range c.lines {
		innerAttrs := strings.Split(line, "=")
		innerShortCut := innerAttrs[0]
		if param == innerShortCut {
			c.lines = append(c.lines[:index], c.lines[index+1:]...)
			break
		}
	}

	if nlines == len(c.lines) {
		errPrint(fmt.Sprintf(`Shortcut "%s" doesn't exist.`, param))
		return
	}

	// Save the record to file
	c.cfgFile.Seek(0, io.SeekStart)
	c.cfgFile.Truncate(0)
	writer := bufio.NewWriter(c.cfgFile)
	defer writer.Flush()
	for _, line := range c.lines {
		writer.WriteString(line + "\n")
	}
}

func (c *CfgRuntime) GetRecordFromCfg(param string) string {
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

func (c *CfgRuntime) ListRecordsFromCfg() {
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

func (c *CfgRuntime) EditRecord() {
	cfgPath := getCfgFilePath()
	if cfgPath == "" {
		return
	}
	cmd := exec.Command("vim", cfgPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
}
