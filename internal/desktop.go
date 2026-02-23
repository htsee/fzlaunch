package internal

import (
	"fmt"
	"os"
	"path/filepath"
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
	Exec       string
	Args       []string
	IsTerminal bool
	Icon       string
	Comment    string
	Categories []string
}

func ParseDesktop(desktopPath string) (string, DesktopEntry, error) {
	content, err := os.ReadFile(desktopPath)
	if err != nil {
		return "", DesktopEntry{}, fmt.Errorf("cannot read file %v", desktopPath)
	}
	lines := strings.Split(string(content), "\n")
	var appName, appExec, appIcon, appComment string
	var appArgs []string
	var isTerm bool
	var appCategories []string
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
				return "", DesktopEntry{}, err
			}
		case "Name":
			appName = val
		case "Exec":
			appExec, _, _ = strings.Cut(val, "%")
			cmd := strings.Split(appExec, " ")
			appExec, appArgs = cmd[0], cmd[1:]
		case "Icon":
			appIcon = val
		case "Comment":
			appComment = val
		case "Categories":
			appCategories = strings.Split(strings.Trim(val, ";"), ";")
		default:
			continue
		}
	}
	return appName, DesktopEntry{Exec: appExec, Args: appArgs, IsTerminal: isTerm, Icon: appIcon, Comment: appComment, Categories: appCategories}, nil
}

func DesktopEntries() (map[string]DesktopEntry, error) {
	entries := make(map[string]DesktopEntry)

	for _, desktopPath := range GetDesktop() {
		name, entry, err := ParseDesktop(desktopPath)
		if err != nil {
			return nil, err
		}
		if entry.IsTerminal || name == "" || entry.Exec == "" {
			continue
		}
		entries[name] = entry
	}
	return entries, nil
}
