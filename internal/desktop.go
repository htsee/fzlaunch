package internal

import (
	"os"
	"path/filepath"
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
