//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

func createWindowsDelayedDelete(installDir string) {
	batchScript := fmt.Sprintf(`@echo off
timeout /t 2 /nobreak > nul
rd /s /q "%s"
del "%%~f0"
`, installDir)

	tempBat := filepath.Join(os.TempDir(), "sundalang_uninstall.bat")
	os.WriteFile(tempBat, []byte(batchScript), 0644)

	cmd := fmt.Sprintf(`start /min cmd /c "%s"`, tempBat)
	os.WriteFile(filepath.Join(os.TempDir(), "sundalang_run.bat"), []byte(cmd), 0644)

	os.StartProcess("cmd", []string{"/c", cmd}, &os.ProcAttr{})
}

func removeFromWindowsPath(pathToRemove string) {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return
	}
	defer k.Close()

	currentPath, _, err := k.GetStringValue("Path")
	if err != nil {
		return
	}

	newPath := ""
	for _, p := range splitPath(currentPath) {
		if !equalsCaseInsensitive(p, pathToRemove) {
			if newPath == "" {
				newPath = p
			} else {
				newPath = newPath + ";" + p
			}
		}
	}

	k.SetStringValue("Path", newPath)
}

func splitPath(path string) []string {
	var result []string
	current := ""
	for _, c := range path {
		if c == ';' {
			if current != "" {
				result = append(result, current)
				current = ""
			}
		} else {
			current = current + string(c)
		}
	}
	if current != "" {
		result = append(result, current)
	}
	return result
}

func equalsCaseInsensitive(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		ca, cb := a[i], b[i]
		if ca >= 'A' && ca <= 'Z' {
			ca = ca + 32
		}
		if cb >= 'A' && cb <= 'Z' {
			cb = cb + 32
		}
		if ca != cb {
			return false
		}
	}
	return true
}
