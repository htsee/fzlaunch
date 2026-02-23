package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
)

var xdgDataDirs = strings.Split(os.Getenv("XDG_DATA_DIRS"), ":")
var xdgDataHome = strings.Split(os.Getenv("XDG_DATA_HOME"), ":")
var XdgData = append(xdgDataDirs, xdgDataHome...)

func GetDesktop() []string {
	var desktop []string
	for _, dir := range XdgData {
		appDir := dir + "/applications"
		files, err := os.ReadDir(appDir)
		if err != nil {
			continue
		}

		for _, file := range files {
			fileName := filepath.Join(appDir, file.Name())
			if filepath.Ext(fileName) == ".desktop" {
				desktop = append(desktop, fileName)
			}
		}
	}
	return desktop
}

type DesktopEntry struct {
	Name       string
	Exec       string
	Args       []string
	IsTerminal bool
}

func ParseDesktop(desktopPath string) (DesktopEntry, error) {
	content, err := os.ReadFile(desktopPath)
	if err != nil {
		return DesktopEntry{}, fmt.Errorf("cannot read file %v", desktopPath)
	}
	lines := strings.Split(string(content), "\n")
	var appName, appExec string
	var appArgs []string
	var isTerm bool
	for _, line := range lines {
		key, val, found := strings.Cut(line, "=")
		if !found {
			if line == "[Desktop Entry]" {
				continue
			} else {
				break
			}
		}
		switch key {
		case "Terminal":
			isTerm, err = strconv.ParseBool(val)
			if err != nil {
				return DesktopEntry{}, err
			}
		case "Name":
			appName = val
		case "Exec":
			appExec, _, _ = strings.Cut(val, "%")
			cmd := strings.Split(appExec, " ")
			appExec, appArgs = cmd[0], cmd[0:]
		default:
			continue
		}
	}
	return DesktopEntry{Name: appName, Exec: appExec, Args: appArgs, IsTerminal: isTerm}, nil
}

func sameName(a, b DesktopEntry) bool {
	return a.Name == b.Name
}

func DesktopEntries() ([]DesktopEntry, error) {
	var entries []DesktopEntry
	for _, desktopPath := range GetDesktop() {
		entry, err := ParseDesktop(desktopPath)
		if err != nil {
			return entries, err
		}
		if entry.IsTerminal || entry.Name == "" || entry.Exec == "" {
			continue
		}
		entries = append(entries, entry)
	}
	entries = slices.CompactFunc(entries, sameName)
	return entries, nil
}
