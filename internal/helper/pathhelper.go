package helper

import (
	"os"
	"path/filepath"
)

func GetExecutableDir() string {
	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		panic(err)
	}
	execDir := filepath.Dir(execPath)
	return execDir
}

func JoinPaths(paths ...string) string {
	return filepath.Join(paths...)
}
