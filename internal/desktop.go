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
	IsTerminal  bool
	Exec        []string
	GenericName string
	Comment     string
	Categories  []string
	Keywords    []string
	Icon        string
}

func ParseDesktop(desktopPath string) (string, DesktopEntry, error) {
	content, err := os.ReadFile(desktopPath)
	if err != nil {
		return "", DesktopEntry{}, fmt.Errorf("cannot read file %q", desktopPath)
	}
	lines := strings.Split(string(content), "\n")
	var isTerm bool
	var appName, appGenericName, appComment, appIcon string
	var appExec, appCategories, appKeywords []string
	for _, line := range lines {
		key, val, found := strings.Cut(line, "=")
		if !found {
			if line == "[Desktop Entry]" || strings.HasPrefix(line, "#") {
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
			cmd, _, _ := strings.Cut(val, "%")
			cmd = strings.TrimSpace(cmd)
			appExec = strings.Split(cmd, " ")
		case "GenericName":
			appGenericName = val
		case "Comment":
			appComment = val
		case "Categories":
			appCategories = strings.Split(strings.Trim(val, ";"), ";")
		case "Keywords":
			appKeywords = strings.Split(strings.Trim(val, ";"), ";")
		case "Icon":
			appIcon = val
		}
	}
	return appName, DesktopEntry{
		IsTerminal:  isTerm,
		Exec:        appExec,
		GenericName: appGenericName,
		Comment:     appComment,
		Categories:  appCategories,
		Keywords:    appKeywords,
		Icon:        appIcon,
	}, nil
}

func DesktopEntries() (map[string]DesktopEntry, error) {
	entries := make(map[string]DesktopEntry)

	for _, desktopPath := range GetDesktop() {
		name, entry, err := ParseDesktop(desktopPath)
		if err != nil {
			return nil, err
		}
		if entry.IsTerminal || name == "" || len(entry.Exec) == 0 {
			continue
		}
		entries[name] = entry
	}
	return entries, nil
}
