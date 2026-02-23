package internal

import (
	"os"
	"strings"
)

var xdgDataDirs = strings.Split(os.Getenv("XDG_DATA_DIRS"), ":")
var xdgDataHome = strings.Split(os.Getenv("XDG_DATA_HOME"), ":")
var XdgData = append(xdgDataDirs, xdgDataHome...)
