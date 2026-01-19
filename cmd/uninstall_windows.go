//go:build windows
// +build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"golang.org/x/sys/windows/registry"
)

func createWindowsDelayedDelete(installDir string) {
	// Robust batch script: loop until file is unlocked (exe exits), then delete.
	// We specifically try to delete the known exe location first to ensure lock is gone,
	// then remove the whole directory.
	exePath := filepath.Join(installDir, "bin", "sundalang.exe")
	batchScript := fmt.Sprintf(`@echo off
:loop
timeout /t 1 /nobreak > nul
del /f /q "%s"
if exist "%s" goto loop
rd /s /q "%s"
del "%%~f0"
`, exePath, exePath, installDir)

	tempBat := filepath.Join(os.TempDir(), "sundalang_uninstall.bat")
	os.WriteFile(tempBat, []byte(batchScript), 0644)

	// Use exec.Command with SysProcAttr to hide the window completely
	c := exec.Command("cmd", "/C", tempBat)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	c.Start()
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
