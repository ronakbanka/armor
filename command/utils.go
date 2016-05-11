package command

import (
	"os"
	"runtime"
)

//BaseDir for armor config path
var (
	BaseDir = os.Getenv("ARMOR_CONFIG_PATH")
)

//GetHomeDir returns the home directory
func GetHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

//GetBaseDir returns BaseDir
func GetBaseDir() string {
	if BaseDir == "" {
		BaseDir = GetHomeDir()
	}
	return BaseDir
}
